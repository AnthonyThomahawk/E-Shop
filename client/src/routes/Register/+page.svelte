<script lang="ts">
    import {registerUser} from "../../lib/Auth";
    import {onMount} from "svelte";
    import {getLocalStorage} from "../../lib/LocalStorage";

    let email = "";
    let password = "";

    let notification = "";
    let notification_color = "black";

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

    const validateEmail = (email) => {
        return email.match(
            /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        );
    };

    async function register() {
        const blankFields = getBlankFields();

        if (blankFields.length == 0) {
            if (password.length >= 8) {
                if (validateEmail(email)) {
                    try {
                        await registerUser(email, password);

                        notification = "E-mail : " + email + " registered, now you can login using login page";
                        notification_color = "green";
                    }
                    catch (error) {
                        notification = "Registration failed, e-mail already exists";
                        notification_color = "red";
                    }
                } else {
                    notification = "Registration failed, " + email + " is not a valid email address.";
                    notification_color = "red";
                }
            } else {
                notification = "Registration failed, password must be 8 characters or more."
                notification_color = "red";
            }
        } else {
            notification = "Registration failed because the field(s) : "
            for (let i = 0; i < blankFields.length; i++) {
                notification += blankFields[i];
                if (i != blankFields.length - 1) {
                    notification += ',';
                }
                notification += ' ';
            }
            notification += " are required."
            notification_color = "red";
        }
    }

    onMount(async () => {
        document.getElementById("email-input")
            .addEventListener("keyup", function(event) {
                event.preventDefault();
                if (event.keyCode === 13) {
                    register();
                }
            });

        document.getElementById("password-input")
            .addEventListener("keyup", function(event) {
                event.preventDefault();
                if (event.keyCode === 13) {
                    register();
                }
            });
    });
</script>

<div class="centered-div">
    <h1>Create a user account</h1>
    <div style="padding-top: 25px;"></div>

    <div class="border-div">
        <label for="email-input">E-mail</label>
        <div>
            <div class="pad-div">
                <img src="/assets/mailicon.webp" alt="" style="border-radius: 5%; width:35px; height:35px; padding-bottom: 10px"/>
            </div>
            <input id="email-input" bind:value={email}>
        </div>

        <label for="password-input">Password</label>
        <div>
            <div class="pad-div">
                <img src="/assets/keyicon.png" alt="" style="border-radius: 5%; width:35px; height:35px; padding-bottom: 10px"/>
            </div>
            <input id="password-input" type="password" bind:value={password}>
        </div>

    </div>

    <div style="padding-top:25px"></div>

    <button on:click={register} style="width: 8vw; height: 3vw;">Register user</button>

    <div style="padding-top:25px"></div>

    <label style="color: {notification_color}">{notification}</label>
    {#if notification_color === "green"}
        <label style="color: green"><a href="/Login">Go to user login page</a></label>
    {/if}
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
        font-size: 1.2vw;
        padding-left: 35px;
    }
    input {
        height:25px;
    }
</style>