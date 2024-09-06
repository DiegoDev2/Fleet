package main

// Hatch - Modern, extensible Python project management
// Homepage: https://hatch.pypa.io/latest/

import (
	"fmt"
	
	"os/exec"
)

func installHatch() {
	// Método 1: Descargar y extraer .tar.gz
	hatch_tar_url := "https://files.pythonhosted.org/packages/fd/40/dbf99436f18bd8b820d5690dff5a534092e1456ba74a87699118b0221417/hatch-1.12.0.tar.gz"
	hatch_cmd_tar := exec.Command("curl", "-L", hatch_tar_url, "-o", "package.tar.gz")
	err := hatch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hatch_zip_url := "https://files.pythonhosted.org/packages/fd/40/dbf99436f18bd8b820d5690dff5a534092e1456ba74a87699118b0221417/hatch-1.12.0.zip"
	hatch_cmd_zip := exec.Command("curl", "-L", hatch_zip_url, "-o", "package.zip")
	err = hatch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hatch_bin_url := "https://files.pythonhosted.org/packages/fd/40/dbf99436f18bd8b820d5690dff5a534092e1456ba74a87699118b0221417/hatch-1.12.0.bin"
	hatch_cmd_bin := exec.Command("curl", "-L", hatch_bin_url, "-o", "binary.bin")
	err = hatch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hatch_src_url := "https://files.pythonhosted.org/packages/fd/40/dbf99436f18bd8b820d5690dff5a534092e1456ba74a87699118b0221417/hatch-1.12.0.src.tar.gz"
	hatch_cmd_src := exec.Command("curl", "-L", hatch_src_url, "-o", "source.tar.gz")
	err = hatch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hatch_cmd_direct := exec.Command("./binary")
	err = hatch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: uv")
	exec.Command("latte", "install", "uv").Run()
	fmt.Println("Instalando dependencia: cffi")
	exec.Command("latte", "install", "cffi").Run()
	fmt.Println("Instalando dependencia: pycparser")
	exec.Command("latte", "install", "pycparser").Run()
}
