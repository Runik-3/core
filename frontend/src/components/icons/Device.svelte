<script>
  export let color = "black";
  export let size = "16px";
  export let hasPlug = false;
</script>

<svg
  width={size}
  viewBox="0 0 129 178"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    fill-rule="evenodd"
    clip-rule="evenodd"
    d="M13.1633 0C5.8934 0 0 5.89337 0 13.1632V133.468C0 140.738 5.8934 146.632 13.1633 146.632H115.837C123.107 146.632 129 140.738 129 133.468V13.1633C129 5.8934 123.107 0 115.837 0H13.1633ZM11.1073 8.21973C9.65329 8.21973 8.47461 9.3984 8.47461 10.8524V124.931C8.47461 126.385 9.65329 127.563 11.1073 127.563H118.207C119.661 127.563 120.84 126.385 120.84 124.931V10.8524C120.84 9.39841 119.661 8.21973 118.207 8.21973H11.1073Z"
    fill={color}
  />
  {#if hasPlug}
    <rect
      x="60.5508"
      y="153.748"
      width="7.89796"
      height="13.1633"
      fill={color}
    />
    <path
      d="M63.1836 166.911H65.8162V177.442H63.1836V169.873V166.911Z"
      fill="url(#paint0_linear_653_337)"
    />
  {/if}
  <defs>
    <linearGradient
      id="paint0_linear_653_337"
      x1="64.4999"
      y1="166.911"
      x2="64.4999"
      y2="177.442"
      gradientUnits="userSpaceOnUse"
    >
      <stop stop-color={color} />
      <stop offset="1" stop-color="#A3A3A3" stop-opacity="0" />
      <stop offset="1" stop-color={color} stop-opacity="0" />
    </linearGradient>
  </defs>
</svg>
