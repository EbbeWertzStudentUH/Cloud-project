<div class="container mx-auto p-6">
  <h1 class="text-4xl font-bold text-center mb-6">Development Page</h1>
  <h2 class="text-2xl font-semibold mb-4">Proxy Forwards</h2>
  <table class="table-auto w-full bg-slate-800 shadow-md rounded-lg overflow-hidden">
    <thead>
      <tr class="bg-emerald-900 text-gray-600 uppercase text-sm leading-normal">
        <th class="py-3 px-6 text-left">Proxy Path</th>
        <th class="py-3 px-6 text-center">Forwarded Port</th>
      </tr>
    </thead>
    <tbody class="text-gray-700 text-sm">
      <tr class="border-b border-gray-200">
        <td class="py-3 px-6 text-left">
          <a href="/proxy/soap" class="text-blue-600 hover:underline">/proxy/soap</a>
        </td>
        <td class="py-3 px-6 text-center">8082</td>
      </tr>
      <tr class="border-b border-gray-200">
        <td class="py-3 px-6 text-left">
          <a href="/proxy/mongo" class="text-blue-600 hover:underline">/proxy/mongo</a>
        </td>
        <td class="py-3 px-6 text-center">8081</td>
      </tr>
      <tr class="border-b border-gray-200">
        <td class="py-3 px-6 text-left">
          <a href="proxy/mysql?server=db-mysql&username=mysql_root&db=cloud_project_db" class="text-blue-600 hover:underline">/proxy/mysql</a>
        </td>
        <td class="py-3 px-6 text-center">8080</td>
      </tr>
    </tbody>
  </table>
</div>
