# Barcode Generator Application

## Overview
The **Barcode Generator** application allows users to easily generate various types of barcodes and save them as PNG images. This user-friendly application supports multiple barcode formats, including **QR Code**, **Code 128**, **Datamatrix**, and **PDF 417**.

Whether you're looking to create a QR code for a URL, generate a unique barcode for inventory management, or create any other type of barcode for business or personal use, this application provides an easy solution.

## Features
- **Generate Barcodes**: Supports four types of barcodes:
  - QR Code
  - Code 128
  - Datamatrix
  - PDF 417
- **Input Custom Text**: Users can input any text or data to generate a corresponding barcode.
- **Save as PNG**: Once the barcode is generated, users can save it to their computer as a PNG file.
- **Error Handling**: The app displays errors directly within the UI, ensuring that users know exactly what went wrong without needing access to the console.
- **Simple UI**: Easy-to-use interface with minimal setup, perfect for both beginners and advanced users.

## Requirements
- **Go** version 1.18 or higher.
- **Fyne** library for building cross-platform GUI applications.

## Installation

### Building the App
1. Clone this repository:
   ```bash
   git clone https://github.com/LewdLillyVT/goqr.git
   ```

2. Navigate into the project folder:
   ```bash
   cd barcode-generator
   ```

3. Install dependencies:
   ```bash
   go get fyne.io/fyne/v2
   go get github.com/boombuler/barcode
   go get github.com/sqweek/dialog
   ```

4. Build the app:
   ```bash
   go build
   ```

5. Run the application:
   ```bash
   ./barcode-generator
   ```

### Running Pre-built Binary
If you're downloading the pre-built binary for your platform, just run the executable after download. Ensure you have the appropriate permissions to execute the file.

## How to Use the App

1. **Enter Barcode Content**: In the input field, enter the text or data that you want to encode into a barcode (e.g., a URL, product code, etc.).
2. **Select Barcode Type**: From the dropdown menu, select the type of barcode you want to generate:
   - **QR Code**: For URL or plain text.
   - **Code 128**: Common for product barcodes.
   - **Datamatrix**: A 2D barcode used for more compact data.
   - **PDF 417**: A 2D barcode often used in government IDs and transport tickets.
3. **Generate Barcode**: Click the "Generate" button to create the barcode. If any error occurs (such as empty input or unsupported barcode type), an error message will be displayed on the screen.
4. **Save Barcode**: After the barcode is generated, click the "Save" button to save the barcode image to your computer. You can choose the file location and name.

## Error Handling
The app displays error messages directly within the app's UI, allowing the user to quickly resolve issues without needing access to the console. Error messages will appear at the bottom of the app interface.

## Screenshots

### Example of the App Interface:
![App Interface](https://cdn.hyrule.pics/feb12619e.png)

### Generated QR Code:
![Generated QR Code](https://cdn.hyrule.pics/e1e1f9c64.png)
