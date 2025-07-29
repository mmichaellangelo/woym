<head>
    <title>What's on your mind?</title>
</head>

<div id="main_frame">
    <h1>What's on your mind?</h1>
    <form method="post">
        <div id="textbox" contenteditable="true" role="textbox"></div>
        <button type="submit">That's it</button>
    </form>
</div>

<style>
    #textbox {
        border: 1px solid black;
        box-sizing: border-box;
        padding: 0.5rem;
        margin-bottom: 1rem;
        width: 25rem;
        word-wrap: break-word;
    }

    #main_frame {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: max-content;
        margin-left: auto;
        margin-right: auto;
    }

    form {
        display: flex;
        flex-direction: column;
        width: 100%;
    }

    button {
        width: fit-content;
    }
</style>