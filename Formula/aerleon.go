package main

// Aerleon - Generate firewall configs for multiple firewall platforms
// Homepage: https://aerleon.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installAerleon() {
	// Método 1: Descargar y extraer .tar.gz
	aerleon_tar_url := "https://files.pythonhosted.org/packages/ca/bd/87869c1cb33a2b4d269c6f66056c44453e643925731cb85e6861d1121be8/aerleon-1.9.0.tar.gz"
	aerleon_cmd_tar := exec.Command("curl", "-L", aerleon_tar_url, "-o", "package.tar.gz")
	err := aerleon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aerleon_zip_url := "https://files.pythonhosted.org/packages/ca/bd/87869c1cb33a2b4d269c6f66056c44453e643925731cb85e6861d1121be8/aerleon-1.9.0.zip"
	aerleon_cmd_zip := exec.Command("curl", "-L", aerleon_zip_url, "-o", "package.zip")
	err = aerleon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aerleon_bin_url := "https://files.pythonhosted.org/packages/ca/bd/87869c1cb33a2b4d269c6f66056c44453e643925731cb85e6861d1121be8/aerleon-1.9.0.bin"
	aerleon_cmd_bin := exec.Command("curl", "-L", aerleon_bin_url, "-o", "binary.bin")
	err = aerleon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aerleon_src_url := "https://files.pythonhosted.org/packages/ca/bd/87869c1cb33a2b4d269c6f66056c44453e643925731cb85e6861d1121be8/aerleon-1.9.0.src.tar.gz"
	aerleon_cmd_src := exec.Command("curl", "-L", aerleon_src_url, "-o", "source.tar.gz")
	err = aerleon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aerleon_cmd_direct := exec.Command("./binary")
	err = aerleon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
