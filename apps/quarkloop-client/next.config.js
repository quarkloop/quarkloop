/** @type {import('next').NextConfig} */

const withBundleAnalyzer = require("@next/bundle-analyzer")({
  enabled: process.env.ANALYZE === "true",
});

// https://github.com/react-dnd/react-dnd/issues/3416#issuecomment-1126159287
// const withTM = require("next-transpile-modules")([
//   "react-dnd",
//   "dnd-core",
//   "@react-dnd/invariant",
//   "@react-dnd/asap",
//   "@react-dnd/shallowequal",
// ]);
//const withTM = require("next-transpile-modules")(["react-dnd"]);

const nextConfig = withBundleAnalyzer({
  reactStrictMode: true,
  swcMinify: true,
  experimental: {
    //typedRoutes: true,
    esmExternals: true,
  },
  transpilePackages: [
    "@quarkloop/types",
    "@quarkloop/lib",
    "@quarkloop/components",
    "@quarkloop/apps",
    "@quarkloop/ui",
  ],
  images: {
    domains: ["i.imgur.com"],
    deviceSizes: [82, 110, 140, 640, 750, 828, 1080, 1200, 1920, 2048, 3840],
    imageSizes: [16, 32, 48, 64],
  },
  //assetPrefix: 'http://localhost:3000',
  async rewrites() {
    return {
      beforeFiles: [
        // {
        //   source: "/docs/:path*",
        //   destination: "http://localhost:3001/docs/:path*",
        // },
        // {
        //   source: "/",
        //   has: [
        //     {
        //       type: "cookie",
        //       key: "next-auth.session-token",
        //     },
        //   ],
        //   destination: "/user",
        // },
        // {
        //   source: "/",
        //   missing: [
        //     {
        //       type: "cookie",
        //       key: "next-auth.session-token",
        //     },
        //   ],
        //   destination: "/signin",
        // },
      ],
    };
  },

  // async redirects() {
  //   return [
  //     {
  //       source: "/signin",
  //       has: [
  //         {
  //           type: "cookie",
  //           key: "next-auth.session-token",
  //         },
  //       ],
  //       destination: "/",
  //       permanent: false,
  //     },
  //   ];
  // },
});

module.exports = nextConfig;
