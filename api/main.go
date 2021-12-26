package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	template := `
	<html lang="en" class="">
  <head>
    <meta charset="UTF-8" />
    <title>Incremental Static Regen</title>

    <meta name="robots" content="noindex" />

    <style>
      @import url("https://fonts.googleapis.com/css2?family=Bungee+Shade&family=Fredoka+One&family=PT+Mono&family=Pushster&display=swap");

      html,
      body {
        margin: 0;
        padding: 0;
      }

      .display {
        display: flex;
        justify-content: center;
        color: white;
        height: 100vh;
        width: 100vw;
        background: #08c9f9;
        background: linear-gradient(to bottom right, #08c9f9, #f00d83);
      }

      .modal {
        text-align: center;
        margin: auto;
      }

      h1 {
        font-size: 5rem;
        font-family: "Pushster", cursive;
      }

      .caption {
        font-family: "Fredoka One", cursive;
        font-size: 2rem;
      }

      strong {
        font-family: "Bungee Shade", cursive;
      }

      .clock {
        font-size: 6rem;
        font-family: "PT Mono", monospace;
        padding-top: 2rem;
      }
    </style>
  </head>

  <body>
    <div class="display">
      <div class="modal">
        <h1>Incremental Static Regen</h1>
        <div class="caption">
          A Demo written in <strong>Go</strong> and hosted on
          <strong>Vercel</strong>
        </div>
        <div class="clock">%s</div>
      </div>
    </div>
  </body>
</html>
`
	currentTime := time.Now().Format("KK:mm:ss a")

	output := fmt.Sprintf(template, currentTime)

	fmt.Fprintf(w, output)
}
