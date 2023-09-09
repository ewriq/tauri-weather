<script>
  // @ts-ignore
  import axios from "axios";

  let sehir = "";
  let error = "";
  let data = {};
  let name = "";
  let temp = ""

  async function get() {
    const form = {
      sehir,
    };
    try {
      const response = await axios.get(`http://localhost:3000/api/${form.sehir}`);
      data = response.data;
      name = response.data.name;
      temp = response.data.main.temp;
      console.log("saa", response.data);
    } catch (err) {
      console.error("saa", err);
      error = err;
    }
  }
</script>


<div class="sağ mt-10">
  {#if error}
  <div class="w-80">
    <div class="flex items-center p-4 mb-4 text-sm text-red-800 border border-red-300 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400 dark:border-red-800" role="alert">
      <svg class="flex-shrink-0 inline w-4 h-4 mr-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
        <path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z"/>
      </svg>
      <div>
        <span class="font-medium">Hata !</span>   {error}
      </div>
    </div>

  </div>
  {/if}

  <form on:submit|preventDefault={get}>
    <div class="mb-6 w-80">
      <label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Ülke </label>
      <input placeholder="Ülke gir knk" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" bind:value={sehir} />
    </div>
    <button class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800" type="submit">sex !</button>
  </form>
 
  {#if data.name}
  <div class="mt-10">
    <a class="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700">
      <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{name} Hava durumu !</h5>
      <p class="font-normal text-gray-700 dark:text-gray-400">{name} Şu anda {temp} derece sıcaklıga sahip</p>
  </a>
  </div>
  {:else}
  <div class="mt-10">
    <a class="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-slate-700 dark:border-gray-700 dark:hover:bg-slate-700">
      <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">Sanırım türkiyede böyle bir yer yok yada bir mesaj yok !</h5>
  </a>
  </div>
  {/if}

</div>