<script lang="ts">
    import DynamicImage from "../../lib/DynamicImage.svelte";
    import UserIcon from "../../assets/usericon.png"
    import KeyIcon from "../../assets/keyicon.png"

    let username = "";
    let password = "";

    let notification = "";
    let notification_color = "black";

    function getBlankFields() {
        let blankFields : Array<string> = [];

        if (username === "") {
            blankFields.push("Username");
        }
        if (password === "") {
            blankFields.push("Password");
        }

        return blankFields;
    }

    function registerUser() {
        const blankFields = getBlankFields();

        if (blankFields.length == 0) {
            notification = "Login successful!";
            notification_color = "green";
        }
        else {
            notification = "Login failed because the field(s) : "
            for (let i = 0; i < blankFields.length; i++){
                notification += blankFields[i];
                if (i != blankFields.length-1) {
                    notification += ',';
                }
                notification += ' ';
            }
            notification += " are required."
            notification_color = "red";
        }
    }
</script>

<div class="centered-div">
    <h1>User login</h1>
    <div style="padding-top: 25px;"></div>

    <div class="border-div">
        <label for="username-input">Username</label>
        <div>
            <div style="transform: translate(0,15%)">
                <DynamicImage imageHeight="35" imageWidth="35" imagePath="{UserIcon}"/>
            </div>
            <input id="username-input" bind:value={username}>
        </div>

        <label for="password-input">Password</label>
        <div>
            <div style="transform: translate(0,15%)">
                <DynamicImage imageHeight="35" imageWidth="35" imagePath="{KeyIcon}"/>
            </div>
            <input id="password-input" type="password" bind:value={password}>
        </div>
    </div>

    <div style="padding-top:25px"></div>

    <button on:click={registerUser} style="width: 200px; height: 55px;">Log in</button>

    <div style="padding-top:25px"></div>

    <label style="color: {notification_color}">{notification}</label>
</div>


<style>
    .border-div {
        padding: 35px;
        border-radius:3%;
        border: thin solid black;
        flex-direction: column;
        align-items: center;
        justify-items: center;
    }
    .centered-div {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }
    div {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
    }
    label {
        font-size: 20px;
        padding-left: 35px;
    }
    input {
        height:25px;
    }
</style>