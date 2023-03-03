<script>
    import {onMount} from "svelte";

    export let imagePath = ""; // pass either a path to a local image, or a URL in the WWW
    export let imageLink = "";
    export let imageHeight = 128;
    export let imageWidth = 128;
    let imageSrc = "";
    async function loadImage() {
        if (imagePath !== "")
        {
            const module = await import(imagePath);
            imageSrc = module.default;
        }
        else if (imageLink !== "")
        {
            const img = new Image();
            img.src = imageLink;
            imageSrc = img.src;
        }
    }

    onMount(loadImage);
</script>
<img src={imageSrc} alt="" style="border-radius: 5%; width:{imageWidth}px; height:{imageHeight}px; padding-bottom: 10px"/>