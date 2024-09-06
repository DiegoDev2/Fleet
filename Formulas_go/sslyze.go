package main

// Sslyze - SSL scanner
// Homepage: https://github.com/nabla-c0d3/sslyze

import (
	"fmt"
	
	"os/exec"
)

func installSslyze() {
	// Método 1: Descargar y extraer .tar.gz
	sslyze_tar_url := "https://files.pythonhosted.org/packages/13/00/bacbb04d7d3e0d7db3cedec0b7a450a6ee9543aa4929b020a329f184daae/sslyze-5.1.3.tar.gz"
	sslyze_cmd_tar := exec.Command("curl", "-L", sslyze_tar_url, "-o", "package.tar.gz")
	err := sslyze_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sslyze_zip_url := "https://files.pythonhosted.org/packages/13/00/bacbb04d7d3e0d7db3cedec0b7a450a6ee9543aa4929b020a329f184daae/sslyze-5.1.3.zip"
	sslyze_cmd_zip := exec.Command("curl", "-L", sslyze_zip_url, "-o", "package.zip")
	err = sslyze_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sslyze_bin_url := "https://files.pythonhosted.org/packages/13/00/bacbb04d7d3e0d7db3cedec0b7a450a6ee9543aa4929b020a329f184daae/sslyze-5.1.3.bin"
	sslyze_cmd_bin := exec.Command("curl", "-L", sslyze_bin_url, "-o", "binary.bin")
	err = sslyze_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sslyze_src_url := "https://files.pythonhosted.org/packages/13/00/bacbb04d7d3e0d7db3cedec0b7a450a6ee9543aa4929b020a329f184daae/sslyze-5.1.3.src.tar.gz"
	sslyze_cmd_src := exec.Command("curl", "-L", sslyze_src_url, "-o", "source.tar.gz")
	err = sslyze_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sslyze_cmd_direct := exec.Command("./binary")
	err = sslyze_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pyinvoke")
exec.Command("latte", "install", "pyinvoke")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@1.1")
exec.Command("latte", "install", "openssl@1.1")
	fmt.Println("Instalando dependencia: pycparser")
exec.Command("latte", "install", "pycparser")
	fmt.Println("Instalando dependencia: python-typing-extensions")
exec.Command("latte", "install", "python-typing-extensions")
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
}
