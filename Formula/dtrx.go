package main

// Dtrx - Intelligent archive extraction
// Homepage: https://pypi.org/project/dtrx/

import (
	"fmt"
	
	"os/exec"
)

func installDtrx() {
	// Método 1: Descargar y extraer .tar.gz
	dtrx_tar_url := "https://files.pythonhosted.org/packages/b7/e6/204294b57be7bb5072c217a1c3ddd5acf9b60b006c215e13e11121c04108/dtrx-8.5.3.tar.gz"
	dtrx_cmd_tar := exec.Command("curl", "-L", dtrx_tar_url, "-o", "package.tar.gz")
	err := dtrx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dtrx_zip_url := "https://files.pythonhosted.org/packages/b7/e6/204294b57be7bb5072c217a1c3ddd5acf9b60b006c215e13e11121c04108/dtrx-8.5.3.zip"
	dtrx_cmd_zip := exec.Command("curl", "-L", dtrx_zip_url, "-o", "package.zip")
	err = dtrx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dtrx_bin_url := "https://files.pythonhosted.org/packages/b7/e6/204294b57be7bb5072c217a1c3ddd5acf9b60b006c215e13e11121c04108/dtrx-8.5.3.bin"
	dtrx_cmd_bin := exec.Command("curl", "-L", dtrx_bin_url, "-o", "binary.bin")
	err = dtrx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dtrx_src_url := "https://files.pythonhosted.org/packages/b7/e6/204294b57be7bb5072c217a1c3ddd5acf9b60b006c215e13e11121c04108/dtrx-8.5.3.src.tar.gz"
	dtrx_cmd_src := exec.Command("curl", "-L", dtrx_src_url, "-o", "source.tar.gz")
	err = dtrx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dtrx_cmd_direct := exec.Command("./binary")
	err = dtrx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: p7zip")
	exec.Command("latte", "install", "p7zip").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
