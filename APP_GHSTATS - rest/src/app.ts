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
    const date = new Date();
    const dayOfWeek = date.getUTCDay(); // 0 = Sunday ... 6 = Saturday

    const sundayUTC = new Date(Date.UTC(
      date.getUTCFullYear(),
      date.getUTCMonth(),
      date.getUTCDate() - dayOfWeek,
      0, 0, 0, 0
  ));

    return Math.floor(sundayUTC.getTime() / 1000); // to Unix
};

const fetchContributorStats = async (repoName: string) => {
  const apiUrl = `https://api.github.com/repos/${repoName}/stats/contributors`;
  let maxRetries = 5;
  let retryDelay = 3000;

  for (let attempt = 0; attempt < maxRetries; attempt++) {
    try {
      const response = await axios.get(apiUrl, {
        headers: {
          "User-Agent": "Node.js-App",
          Authorization: `Bearer ${process.env.GH_BEARER_TOKEN}`,
        },
      });

      if (response.status === 200) {
        return response.data;
      } else if (response.status === 202) {
        console.log("Data is being prepared. Retrying...");
        await new Promise((resolve) => setTimeout(resolve, retryDelay));
      }
    } catch (error) {
      console.error("Error fetching stats:");
      throw error;
    }
  }
  throw new Error("Max retries reached. The data is not ready yet.");
};

app.post("/github-stats/", async (req: Request, res: Response): Promise<any> => {
  const { githubUrl } = req.body;

  const repoName = extractRepoName(githubUrl);
  if (!repoName) {
    return res.status(400).json({ error: "Invalid GitHub repository URL" });
  }

  try {
    const contributors = await fetchContributorStats(repoName);
    console.log(contributors)
    const lastWeekStart = getLastWeekStart();
    console.log("last week: ", lastWeekStart)

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
        (week: any) => week.w == lastWeekStart
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
