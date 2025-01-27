import express, { Request, Response } from "express";
import axios from "axios";
import dotenv from "dotenv";
import cors from "cors";


dotenv.config();


const app = express();
const PORT = process.env.LISTEN_PORT;

app.use(cors());

app.use(express.json());


const extractRepoName = (ghUrl: string): string | null => {
  const match = ghUrl.match(/github\.com\/([^/]+\/[^/]+)/);
  return match ? match[1] : null;
};

const getLastWeekStart = (): number => {
  const currentTimestamp = Math.floor(Date.now() / 1000);
  const secondsInWeek = 7 * 24 * 60 * 60;
  return currentTimestamp - (currentTimestamp % secondsInWeek);
};

app.post("/github-stats/", async (req: Request, res: Response): Promise<any> => {
  const { githubUrl } = req.body;

  const repoName = extractRepoName(githubUrl);
  if (!repoName) {
    return res.status(400).json({ error: "Invalid GitHub repository URL" });
  }

  const apiUrl = `https://api.github.com/repos/${repoName}/stats/contributors`;
  try {
    const response = await axios.get(apiUrl, {
      headers: {
        "User-Agent": "Node.js-App" ,
        Authorization: `Bearer ${process.env.GH_BEARER_TOKEN}`,
      },
    });

    
    const contributors = response.data;
    const lastWeekStart = getLastWeekStart();

    const stats = contributors.map((contributor: any) => {
      const username = contributor.author.login;
      const totalCommits = contributor.total;
      const totalAdditions = contributor.weeks.reduce(
        (sum: number, week: any) => sum + week.a,
        0
      );
      const totalDeletions = contributor.weeks.reduce(
        (sum: number, week: any) => sum + week.d,
        0
      );

      const lastWeek = contributor.weeks.find(
        (week: any) => week.w === lastWeekStart
      ) || { a: 0, d: 0, c: 0 };

      return {
        username,
        totalCommits,
        totalAdditions,
        totalDeletions,
        lastWeekCommits: lastWeek.c,
        lastWeekAdditions: lastWeek.a,
        lastWeekDeletions: lastWeek.d,
      };
    });

    stats.sort((a: any, b: any) => b.totalCommits - a.totalCommits);

    res.json(stats);
  } catch (error: any) {
    console.log(error)
    res
      .status(error.response?.status || 500)
      .json({ error: error.response?.data || "GitHub API error" });
  }
});

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
