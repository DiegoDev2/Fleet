package main

// Vunnel - Tool for collecting vulnerability data from various sources
// Homepage: https://github.com/anchore/vunnel

import (
	"fmt"
	
	"os/exec"
)

func installVunnel() {
	// Método 1: Descargar y extraer .tar.gz
	vunnel_tar_url := "https://files.pythonhosted.org/packages/ba/f0/96d50b5c15307c3463911028ce14eb0575c8793e1c4d21d8113c90a8f13e/vunnel-0.27.0.tar.gz"
	vunnel_cmd_tar := exec.Command("curl", "-L", vunnel_tar_url, "-o", "package.tar.gz")
	err := vunnel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vunnel_zip_url := "https://files.pythonhosted.org/packages/ba/f0/96d50b5c15307c3463911028ce14eb0575c8793e1c4d21d8113c90a8f13e/vunnel-0.27.0.zip"
	vunnel_cmd_zip := exec.Command("curl", "-L", vunnel_zip_url, "-o", "package.zip")
	err = vunnel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vunnel_bin_url := "https://files.pythonhosted.org/packages/ba/f0/96d50b5c15307c3463911028ce14eb0575c8793e1c4d21d8113c90a8f13e/vunnel-0.27.0.bin"
	vunnel_cmd_bin := exec.Command("curl", "-L", vunnel_bin_url, "-o", "binary.bin")
	err = vunnel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vunnel_src_url := "https://files.pythonhosted.org/packages/ba/f0/96d50b5c15307c3463911028ce14eb0575c8793e1c4d21d8113c90a8f13e/vunnel-0.27.0.src.tar.gz"
	vunnel_cmd_src := exec.Command("curl", "-L", vunnel_src_url, "-o", "source.tar.gz")
	err = vunnel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vunnel_cmd_direct := exec.Command("./binary")
	err = vunnel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
