interface ErrorResponse {
  message?: string;
}

async function getErrorMessageFromResponse(
  response: Response,
  defaultMessage: string,
): Promise<string> {
  let errorMessage = defaultMessage;
  try {
    const errorData = (await response.json()) as ErrorResponse;
    errorMessage = errorData.message || errorMessage;
  } catch (e) {
    console.error("Failed to parse error response", e);
  }
  return errorMessage;
}

export { getErrorMessageFromResponse };
