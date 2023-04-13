# DevBook_Web

DevBook is a social network developed in the Go programming language, where users can create their profiles, make posts and follow other users. These three use cases form the core of the application, while DevBook_App is the server-side application responsible for processing and persisting data in a database from requests made by the [DevBook_Web](https://github.com/CarlosCezarDeSouzaGuaraldo/DevBook_Web) client.

## What do you need to install?

1. [Git](https://git-scm.com/) (any version)
2. [GO](https://go.dev/) (any version)
3. [MySQL](https://www.mysql.com/) (any version)
4. [GitHub Desktop](https://desktop.github.com/) (it's not mandatory)

## How to run?

1. Clone the repository [DevBook_App](https://github.com/CarlosCezarDeSouzaGuaraldo/DevBook_App) from branch ```main```
2. Access MySQL and execute the script for creating the database and structuring the tables according to the [/sql/sql.sql](/sql/sql.sql) file of the application.
3. Create at the root of the project one file ```.env``` based on the ```template.env``` file and fill the .env file according to the settings you want to execute.
4. Open the command prompt and navigate to the root directory of the application.
5. Execute the command line ```go run api``` and type ```Enter```.
6. If everything went well, a message like this should appear in your command prompt: ```API running on localhost:5000```

Then, you can access the application in a browser based on the information provided in the environment variables **HOST** and **PORT** in your ```.env``` file.

<div align="center">
 <img src="https://user-images.githubusercontent.com/66181262/231670408-4f0d96ec-eb09-4526-a8dc-c66dae28b655.jpg" />
</div>
