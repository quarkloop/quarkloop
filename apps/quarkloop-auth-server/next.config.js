/** @type {import('next').NextConfig} */

const withBundleAnalyzer = require("@next/bundle-analyzer")({
    enabled: process.env.ANALYZE === "true",
});

const nextConfig = withBundleAnalyzer({
    swcMinify: true,
    experimental: {
        esmExternals: true,
    },
    transpilePackages: [
        "@quarkloop/components",
        "@quarkloop/lib",
        "@quarkloop/types"
    ],
});

module.exports = nextConfig;
