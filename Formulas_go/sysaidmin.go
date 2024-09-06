package main

// Sysaidmin - GPT-powered sysadmin
// Homepage: https://github.com/skorokithakis/sysaidmin

import (
	"fmt"
	
	"os/exec"
)

func installSysaidmin() {
	// Método 1: Descargar y extraer .tar.gz
	sysaidmin_tar_url := "https://files.pythonhosted.org/packages/59/01/4f961856c3cfb32136782b119b3944466df81fb542dc6c214dc3386245f0/sysaidmin-0.2.2.tar.gz"
	sysaidmin_cmd_tar := exec.Command("curl", "-L", sysaidmin_tar_url, "-o", "package.tar.gz")
	err := sysaidmin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sysaidmin_zip_url := "https://files.pythonhosted.org/packages/59/01/4f961856c3cfb32136782b119b3944466df81fb542dc6c214dc3386245f0/sysaidmin-0.2.2.zip"
	sysaidmin_cmd_zip := exec.Command("curl", "-L", sysaidmin_zip_url, "-o", "package.zip")
	err = sysaidmin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sysaidmin_bin_url := "https://files.pythonhosted.org/packages/59/01/4f961856c3cfb32136782b119b3944466df81fb542dc6c214dc3386245f0/sysaidmin-0.2.2.bin"
	sysaidmin_cmd_bin := exec.Command("curl", "-L", sysaidmin_bin_url, "-o", "binary.bin")
	err = sysaidmin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sysaidmin_src_url := "https://files.pythonhosted.org/packages/59/01/4f961856c3cfb32136782b119b3944466df81fb542dc6c214dc3386245f0/sysaidmin-0.2.2.src.tar.gz"
	sysaidmin_cmd_src := exec.Command("curl", "-L", sysaidmin_src_url, "-o", "source.tar.gz")
	err = sysaidmin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sysaidmin_cmd_direct := exec.Command("./binary")
	err = sysaidmin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
