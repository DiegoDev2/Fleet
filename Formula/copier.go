package main

// Copier - Utility for rendering projects templates
// Homepage: https://copier.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installCopier() {
	// Método 1: Descargar y extraer .tar.gz
	copier_tar_url := "https://files.pythonhosted.org/packages/0f/32/56044ea18ea1c6dc9a00714cbaf7c0ff7e5f486adbfefc28f5ca72a14064/copier-9.3.1.tar.gz"
	copier_cmd_tar := exec.Command("curl", "-L", copier_tar_url, "-o", "package.tar.gz")
	err := copier_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	copier_zip_url := "https://files.pythonhosted.org/packages/0f/32/56044ea18ea1c6dc9a00714cbaf7c0ff7e5f486adbfefc28f5ca72a14064/copier-9.3.1.zip"
	copier_cmd_zip := exec.Command("curl", "-L", copier_zip_url, "-o", "package.zip")
	err = copier_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	copier_bin_url := "https://files.pythonhosted.org/packages/0f/32/56044ea18ea1c6dc9a00714cbaf7c0ff7e5f486adbfefc28f5ca72a14064/copier-9.3.1.bin"
	copier_cmd_bin := exec.Command("curl", "-L", copier_bin_url, "-o", "binary.bin")
	err = copier_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	copier_src_url := "https://files.pythonhosted.org/packages/0f/32/56044ea18ea1c6dc9a00714cbaf7c0ff7e5f486adbfefc28f5ca72a14064/copier-9.3.1.src.tar.gz"
	copier_cmd_src := exec.Command("curl", "-L", copier_src_url, "-o", "source.tar.gz")
	err = copier_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	copier_cmd_direct := exec.Command("./binary")
	err = copier_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
