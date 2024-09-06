package main

// GnuBarcode - Convert text strings to printed bars
// Homepage: https://www.gnu.org/software/barcode/

import (
	"fmt"
	
	"os/exec"
)

func installGnuBarcode() {
	// Método 1: Descargar y extraer .tar.gz
	gnubarcode_tar_url := "https://ftp.gnu.org/gnu/barcode/barcode-0.99.tar.gz"
	gnubarcode_cmd_tar := exec.Command("curl", "-L", gnubarcode_tar_url, "-o", "package.tar.gz")
	err := gnubarcode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnubarcode_zip_url := "https://ftp.gnu.org/gnu/barcode/barcode-0.99.zip"
	gnubarcode_cmd_zip := exec.Command("curl", "-L", gnubarcode_zip_url, "-o", "package.zip")
	err = gnubarcode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnubarcode_bin_url := "https://ftp.gnu.org/gnu/barcode/barcode-0.99.bin"
	gnubarcode_cmd_bin := exec.Command("curl", "-L", gnubarcode_bin_url, "-o", "binary.bin")
	err = gnubarcode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnubarcode_src_url := "https://ftp.gnu.org/gnu/barcode/barcode-0.99.src.tar.gz"
	gnubarcode_cmd_src := exec.Command("curl", "-L", gnubarcode_src_url, "-o", "source.tar.gz")
	err = gnubarcode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnubarcode_cmd_direct := exec.Command("./binary")
	err = gnubarcode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
