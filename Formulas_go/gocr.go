package main

// Gocr - Optical Character Recognition (OCR), converts images back to text
// Homepage: https://wasd.urz.uni-magdeburg.de/jschulen/ocr/

import (
	"fmt"
	
	"os/exec"
)

func installGocr() {
	// Método 1: Descargar y extraer .tar.gz
	gocr_tar_url := "https://wasd.urz.uni-magdeburg.de/jschulen/ocr/gocr-0.52.tar.gz"
	gocr_cmd_tar := exec.Command("curl", "-L", gocr_tar_url, "-o", "package.tar.gz")
	err := gocr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gocr_zip_url := "https://wasd.urz.uni-magdeburg.de/jschulen/ocr/gocr-0.52.zip"
	gocr_cmd_zip := exec.Command("curl", "-L", gocr_zip_url, "-o", "package.zip")
	err = gocr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gocr_bin_url := "https://wasd.urz.uni-magdeburg.de/jschulen/ocr/gocr-0.52.bin"
	gocr_cmd_bin := exec.Command("curl", "-L", gocr_bin_url, "-o", "binary.bin")
	err = gocr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gocr_src_url := "https://wasd.urz.uni-magdeburg.de/jschulen/ocr/gocr-0.52.src.tar.gz"
	gocr_cmd_src := exec.Command("curl", "-L", gocr_src_url, "-o", "source.tar.gz")
	err = gocr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gocr_cmd_direct := exec.Command("./binary")
	err = gocr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: netpbm")
exec.Command("latte", "install", "netpbm")
}
