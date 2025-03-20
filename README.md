# Go Heartcore Content Migration

### To run this code
Have [Go](https://go.dev/) installed on your machine. \
At the root of the folder, perform the following steps: \
Run this command: `go run main.go`


### Description of task


1. Create a Heartcore project
2. Create Two languages English(US) & Danish(Denmark)
3. Create the following structure in the project

- Create two folders 📁
   - **Elements**
   - **pages**
- In the **Elements folder 📂** create an **Element type ⚛️** called **Genre** and create two properties:
   - **Index number(Textstring, Mandatory)**
   - **Title(Textstring, Vary by culture)**
- In the **Pages 📁** folder create two **Document Types**:
   - **TV Show**
   - **TV Shows**
- **TV Show** Document Type needs 3 properties:
   - **Show Summary(RTE, Vary by culture)**
   - **Show Genres(Blocklist using Element type: Genre)**
   - **Show Image(Image Media Picker)**
   - **Show Id(Label(String))**
- **TV Shows** Document Type has no properties and has the following settings:
   - List View ✔️
   - Allow as root ✔️
   - Allowed child node types: **TV Show**

- The structure ↓
   - ![image](https://github.com/user-attachments/assets/b4d17bd8-ba88-40ea-8732-592355448cb4)
4. Use the [Heartcore API Docs](https://docs.umbraco.com/umbraco-heartcore/api-documentation/api-documentation) and the [api.tvmaze.com](https://www.tvmaze.com/api) to create a content node that is populated with the TV Shows automatically.
An example of how that could look would be like so:
![Screenshot 2025-03-20 at 12 41 22](https://github.com/user-attachments/assets/4938fa5d-0681-4001-a29a-c4c904c5a8a1)

5.  If a show exists with an ID do the following ↓
- Update Existing Property:
   - If the property already exists in Heartcore but is empty, check if TVMAZE has a value for it. If it does, update the property in Heartcore with the value from TVMAZE.
- Create New Content with Media:
   - If the property doesn't exist in Heartcore, download the images from TVMAZE, save them to the media section in Heartcore, and link them to a new content node you create.
