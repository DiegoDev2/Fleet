package main

// SshMitm - SSH server for security audits and malware analysis
// Homepage: https://docs.ssh-mitm.at

import (
	"fmt"
	
	"os/exec"
)

func installSshMitm() {
	// Método 1: Descargar y extraer .tar.gz
	sshmitm_tar_url := "https://files.pythonhosted.org/packages/65/22/d5a7a153b1f40f31e1a7e15439e4e3a2aad1413a486aa69c2f0be6482295/ssh_mitm-5.0.0.tar.gz"
	sshmitm_cmd_tar := exec.Command("curl", "-L", sshmitm_tar_url, "-o", "package.tar.gz")
	err := sshmitm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sshmitm_zip_url := "https://files.pythonhosted.org/packages/65/22/d5a7a153b1f40f31e1a7e15439e4e3a2aad1413a486aa69c2f0be6482295/ssh_mitm-5.0.0.zip"
	sshmitm_cmd_zip := exec.Command("curl", "-L", sshmitm_zip_url, "-o", "package.zip")
	err = sshmitm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sshmitm_bin_url := "https://files.pythonhosted.org/packages/65/22/d5a7a153b1f40f31e1a7e15439e4e3a2aad1413a486aa69c2f0be6482295/ssh_mitm-5.0.0.bin"
	sshmitm_cmd_bin := exec.Command("curl", "-L", sshmitm_bin_url, "-o", "binary.bin")
	err = sshmitm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sshmitm_src_url := "https://files.pythonhosted.org/packages/65/22/d5a7a153b1f40f31e1a7e15439e4e3a2aad1413a486aa69c2f0be6482295/ssh_mitm-5.0.0.src.tar.gz"
	sshmitm_cmd_src := exec.Command("curl", "-L", sshmitm_src_url, "-o", "source.tar.gz")
	err = sshmitm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sshmitm_cmd_direct := exec.Command("./binary")
	err = sshmitm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
