export function getFaviconUrl(url: string): string {
  try {
    const domain = new URL(url).hostname;
    // Use Google's favicon service
    return `https://www.google.com/s2/favicons?domain=${domain}&sz=64`;
  } catch {
    return '';
  }
}

export function getDomain(url: string): string {
  try {
    return new URL(url).hostname;
  } catch {
    return url;
  }
}
