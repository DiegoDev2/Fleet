package main

// Fail2ban - Scan log files and ban IPs showing malicious signs
// Homepage: https://www.fail2ban.org/

import (
	"fmt"
	
	"os/exec"
)

func installFail2ban() {
	// Método 1: Descargar y extraer .tar.gz
	fail2ban_tar_url := "https://github.com/fail2ban/fail2ban/archive/refs/tags/1.1.0.tar.gz"
	fail2ban_cmd_tar := exec.Command("curl", "-L", fail2ban_tar_url, "-o", "package.tar.gz")
	err := fail2ban_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fail2ban_zip_url := "https://github.com/fail2ban/fail2ban/archive/refs/tags/1.1.0.zip"
	fail2ban_cmd_zip := exec.Command("curl", "-L", fail2ban_zip_url, "-o", "package.zip")
	err = fail2ban_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fail2ban_bin_url := "https://github.com/fail2ban/fail2ban/archive/refs/tags/1.1.0.bin"
	fail2ban_cmd_bin := exec.Command("curl", "-L", fail2ban_bin_url, "-o", "binary.bin")
	err = fail2ban_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fail2ban_src_url := "https://github.com/fail2ban/fail2ban/archive/refs/tags/1.1.0.src.tar.gz"
	fail2ban_cmd_src := exec.Command("curl", "-L", fail2ban_src_url, "-o", "source.tar.gz")
	err = fail2ban_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fail2ban_cmd_direct := exec.Command("./binary")
	err = fail2ban_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
