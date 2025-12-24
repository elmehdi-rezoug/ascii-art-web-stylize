# ASCII Art Web

## Description

A web app that turns your text into cool ascii art.
choose from three banner styles (standard, shadow or thinkertoy), type something, and watch it transform into ascii art. built with go and html.

## Authors

- mohammed amine el bouziani (@melbouzi)
- elmehdi rezoug (@erezzoug)
- aymane bouzerda (@abouzerd)

## Usage

### Running the Server

```bash
git clone https://learn.zone01oujda.ma/git/melbouzi/ascii-art-web
cd ascii-art-web
go run .
```

then open `http://localhost:8080` in your browser.

### Using the App

1. type your text in the textarea
2. pick a banner style (radio buttons)
3. hit submit
4. see your ascii art appear below

## Implementation Details

### How It Works

**the algorithm:**

each banner file has 856 lines, that's 8 lines of art for each printable character ascii 32-126 plus spacing.
when you submit text, here's what happens:

1. we validate your input (only printable characters allowed)
2. load the banner file you chose
3. for each character in your text:
   - calculate which lines to grab: `(character ascii code - 32) × 9 + line_height`
   - pull those 8 lines from the banner
   - stack them horizontally to form your character
4. repeat for all characters, building the final output

**routes:**
- `GET /` → main page with the form
- `POST /ascii-art` → processes your text and returns the art

**error handling:**
- 200 - everything worked
- 400 - bad input (invalid characters or banner)
- 404 - page not found
- 405 - wrong http method
- 500 - something broke on our end

### Project Structure

```
ascii-art-web/
├── asciigenerator/
│   ├── generator.go    
│   └── utils.go       
├── banners/
│   ├── standard.txt   
│   ├── shadow.txt      
│   └── thinkertoy.txt  
├── templates/
│   └── index.html      
├── homePageHandler.go     
├── asciiArtHandler.go      
├── main.go           
└── go.mod             
```

Built with go standard library.