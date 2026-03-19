#!/bin/bash

# Generate sitemap.xml for multi-language documentation
# This script creates sitemaps for SEO optimization

OUTPUT_DIR="doc_build"
BASE_URL="https://lazygophers.github.io/log"
CURRENT_DATE=$(date -u +"%Y-%m-%d")

# Clean old sitemaps
rm -f "$OUTPUT_DIR"/sitemap*.xml

# Create sitemap for zh-CN (default locale)
cat > "$OUTPUT_DIR/sitemap-zh-CN.xml" << EOF
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url><loc>${BASE_URL}/</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>weekly</changefreq><priority>1.0</priority></url>
  <url><loc>${BASE_URL}/API</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>weekly</changefreq><priority>0.9</priority></url>
  <url><loc>${BASE_URL}/CHANGELOG</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>weekly</changefreq><priority>0.8</priority></url>
  <url><loc>${BASE_URL}/CONTRIBUTING</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>monthly</changefreq><priority>0.7</priority></url>
  <url><loc>${BASE_URL}/CODE_OF_CONDUCT</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>monthly</changefreq><priority>0.6</priority></url>
  <url><loc>${BASE_URL}/SECURITY</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>monthly</changefreq><priority>0.6</priority></url>
  <url><loc>${BASE_URL}/LICENSE</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>yearly</changefreq><priority>0.5</priority></url>
</urlset>
EOF

# Create sitemap for en
cat > "$OUTPUT_DIR/sitemap-en.xml" << EOF
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url><loc>${BASE_URL}/en/</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>weekly</changefreq><priority>1.0</priority></url>
  <url><loc>${BASE_URL}/en/API</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>weekly</changefreq><priority>0.9</priority></url>
  <url><loc>${BASE_URL}/en/CHANGELOG</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>weekly</changefreq><priority>0.8</priority></url>
  <url><loc>${BASE_URL}/en/CONTRIBUTING</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>monthly</changefreq><priority>0.7</priority></url>
  <url><loc>${BASE_URL}/en/CODE_OF_CONDUCT</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>monthly</changefreq><priority>0.6</priority></url>
  <url><loc>${BASE_URL}/en/SECURITY</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>monthly</changefreq><priority>0.6</priority></url>
  <url><loc>${BASE_URL}/en/LICENSE</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>yearly</changefreq><priority>0.5</priority></url>
</urlset>
EOF

# Create sitemap for zh-TW
cat > "$OUTPUT_DIR/sitemap-zh-TW.xml" << EOF
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url><loc>${BASE_URL}/zh-TW/</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>weekly</changefreq><priority>1.0</priority></url>
  <url><loc>${BASE_URL}/zh-TW/API</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>weekly</changefreq><priority>0.9</priority></url>
  <url><loc>${BASE_URL}/zh-TW/CHANGELOG</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>weekly</changefreq><priority>0.8</priority></url>
  <url><loc>${BASE_URL}/zh-TW/CONTRIBUTING</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>monthly</changefreq><priority>0.7</priority></url>
  <url><loc>${BASE_URL}/zh-TW/CODE_OF_CONDUCT</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>monthly</changefreq><priority>0.6</priority></url>
  <url><loc>${BASE_URL}/zh-TW/SECURITY</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>monthly</changefreq><priority>0.6</priority></url>
  <url><loc>${BASE_URL}/zh-TW/LICENSE</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>yearly</changefreq><priority>0.5</priority></url>
</urlset>
EOF

# Create sitemap for fr
cat > "$OUTPUT_DIR/sitemap-fr.xml" << EOF
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url><loc>${BASE_URL}/fr/</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>weekly</changefreq><priority>1.0</priority></url>
  <url><loc>${BASE_URL}/fr/API</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>weekly</changefreq><priority>0.9</priority></url>
  <url><loc>${BASE_URL}/fr/CHANGELOG</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>weekly</changefreq><priority>0.8</priority></url>
  <url><loc>${BASE_URL}/fr/CONTRIBUTING</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>monthly</changefreq><priority>0.7</priority></url>
  <url><loc>${BASE_URL}/fr/CODE_OF_CONDUCT</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>monthly</changefreq><priority>0.6</priority></url>
  <url><loc>${BASE_URL}/fr/SECURITY</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>monthly</changefreq><priority>0.6</priority></url>
  <url><loc>${BASE_URL}/fr/LICENSE</loc><lastmod>${CURRENT_DATE}</lastmod><changefreq>yearly</changefreq><priority>0.5</priority></url>
</urlset>
EOF

# Create main sitemap index
cat > "$OUTPUT_DIR/sitemap.xml" << EOF
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <sitemap><loc>https://lazygophers.github.io/log/sitemap-zh-CN.xml</loc></sitemap>
  <sitemap><loc>https://lazygophers.github.io/log/sitemap-en.xml</loc></sitemap>
  <sitemap><loc>https://lazygophers.github.io/log/sitemap-zh-TW.xml</loc></sitemap>
  <sitemap><loc>https://lazygophers.github.io/log/sitemap-fr.xml</loc></sitemap>
</sitemapindex>
EOF

echo "Sitemap generation complete!"
ls -la "$OUTPUT_DIR"/sitemap*.xml
