package main

// Textract - Extract text from various different types of files
// Homepage: https://textract.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installTextract() {
	// Método 1: Descargar y extraer .tar.gz
	textract_tar_url := "https://files.pythonhosted.org/packages/81/9f/dd29fcec368f007d44e51f0273489d5172a6d32ed9c796df5054fbb31c9f/textract-1.6.5.tar.gz"
	textract_cmd_tar := exec.Command("curl", "-L", textract_tar_url, "-o", "package.tar.gz")
	err := textract_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	textract_zip_url := "https://files.pythonhosted.org/packages/81/9f/dd29fcec368f007d44e51f0273489d5172a6d32ed9c796df5054fbb31c9f/textract-1.6.5.zip"
	textract_cmd_zip := exec.Command("curl", "-L", textract_zip_url, "-o", "package.zip")
	err = textract_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	textract_bin_url := "https://files.pythonhosted.org/packages/81/9f/dd29fcec368f007d44e51f0273489d5172a6d32ed9c796df5054fbb31c9f/textract-1.6.5.bin"
	textract_cmd_bin := exec.Command("curl", "-L", textract_bin_url, "-o", "binary.bin")
	err = textract_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	textract_src_url := "https://files.pythonhosted.org/packages/81/9f/dd29fcec368f007d44e51f0273489d5172a6d32ed9c796df5054fbb31c9f/textract-1.6.5.src.tar.gz"
	textract_cmd_src := exec.Command("curl", "-L", textract_src_url, "-o", "source.tar.gz")
	err = textract_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	textract_cmd_direct := exec.Command("./binary")
	err = textract_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: antiword")
	exec.Command("latte", "install", "antiword").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: pillow")
	exec.Command("latte", "install", "pillow").Run()
	fmt.Println("Instalando dependencia: poppler")
	exec.Command("latte", "install", "poppler").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: tesseract")
	exec.Command("latte", "install", "tesseract").Run()
	fmt.Println("Instalando dependencia: unrtf")
	exec.Command("latte", "install", "unrtf").Run()
}
