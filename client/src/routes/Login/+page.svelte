<script lang="ts">
    import DynamicImage from "../../lib/DynamicImage.svelte";
    import MailIcon from "../../assets/mailicon.webp";
    import KeyIcon from "../../assets/keyicon.png";
    import {setLocalStorage} from "../../lib/LocalStorage";
    import {loginUser} from "../../lib/Auth";
    import {goto} from "$app/navigation";

    let email = "";
    let password = "";

    let notification = "";
    let notification_color = "black";

    interface IUserAuth {
        Email: string;
        Token: string;
    }

    let UserAuth : IUserAuth = {} as IUserAuth;

    function getBlankFields() {
        let blankFields : Array<string> = [];

        if (email === "") {
            blankFields.push("E-mail");
        }
        if (password === "") {
            blankFields.push("Password");
        }

        return blankFields;
    }

    async function login() {
        const blankFields = getBlankFields();

        if (blankFields.length == 0) {
            try {
                const res = await loginUser(email, password);

                UserAuth = {
                    Email : res.Email,
                    Token : res.Token
                }

                setLocalStorage("UserData",
                    {
                        Email: UserAuth.Email,
                        Token: UserAuth.Token
                    }
                );

                notification = "Login as "+ UserAuth.Email + " successful!";
                notification_color = "green";

                await new Promise(resolve => setTimeout(resolve, 1000));
                await goto('/');
                location.reload();
            }
            catch (error) {
                notification = "Login failed, wrong e-mail and/or password";
                notification_color = "red";
            }

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
        <label for="email-input">E-mail</label>
        <div>
            <div class="pad-div">
                <DynamicImage imageHeight="35" imageWidth="35" imagePath="{MailIcon}"/>
            </div>
            <input id="email-input" bind:value={email}>
        </div>

        <label for="password-input">Password</label>
        <div>
            <div class="pad-div">
                <DynamicImage imageHeight="35" imageWidth="35" imagePath="{KeyIcon}"/>
            </div>
            <input id="password-input" type="password" bind:value={password}>
        </div>
    </div>

    <div style="padding-top:25px"></div>

    <button on:click={login} style="width: 200px; height: 55px;">Log in</button>

    <div style="padding-top:25px"></div>

    <label style="color: {notification_color}">{notification}</label>
</div>


<style>
    .pad-div {
        margin-right: 5px;
        transform: translate(0,15%)
    }
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