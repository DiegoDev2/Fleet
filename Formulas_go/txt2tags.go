package main

// Txt2tags - Conversion tool to generating several file formats
// Homepage: https://txt2tags.org/

import (
	"fmt"
	
	"os/exec"
)

func installTxt2tags() {
	// Método 1: Descargar y extraer .tar.gz
	txt2tags_tar_url := "https://files.pythonhosted.org/packages/27/17/c9cdebfc86e824e25592a20a8871225dad61b6b6c0101f4a2cb3434890dd/txt2tags-3.9.tar.gz"
	txt2tags_cmd_tar := exec.Command("curl", "-L", txt2tags_tar_url, "-o", "package.tar.gz")
	err := txt2tags_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	txt2tags_zip_url := "https://files.pythonhosted.org/packages/27/17/c9cdebfc86e824e25592a20a8871225dad61b6b6c0101f4a2cb3434890dd/txt2tags-3.9.zip"
	txt2tags_cmd_zip := exec.Command("curl", "-L", txt2tags_zip_url, "-o", "package.zip")
	err = txt2tags_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	txt2tags_bin_url := "https://files.pythonhosted.org/packages/27/17/c9cdebfc86e824e25592a20a8871225dad61b6b6c0101f4a2cb3434890dd/txt2tags-3.9.bin"
	txt2tags_cmd_bin := exec.Command("curl", "-L", txt2tags_bin_url, "-o", "binary.bin")
	err = txt2tags_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	txt2tags_src_url := "https://files.pythonhosted.org/packages/27/17/c9cdebfc86e824e25592a20a8871225dad61b6b6c0101f4a2cb3434890dd/txt2tags-3.9.src.tar.gz"
	txt2tags_cmd_src := exec.Command("curl", "-L", txt2tags_src_url, "-o", "source.tar.gz")
	err = txt2tags_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	txt2tags_cmd_direct := exec.Command("./binary")
	err = txt2tags_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
