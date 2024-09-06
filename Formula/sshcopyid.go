package main

// SshCopyId - Add a public key to a remote machine's authorized_keys file
// Homepage: https://www.openssh.com/

import (
	"fmt"
	
	"os/exec"
)

func installSshCopyId() {
	// Método 1: Descargar y extraer .tar.gz
	sshcopyid_tar_url := "https://cdn.openbsd.org/pub/OpenBSD/OpenSSH/portable/openssh-9.8p1.tar.gz"
	sshcopyid_cmd_tar := exec.Command("curl", "-L", sshcopyid_tar_url, "-o", "package.tar.gz")
	err := sshcopyid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sshcopyid_zip_url := "https://cdn.openbsd.org/pub/OpenBSD/OpenSSH/portable/openssh-9.8p1.zip"
	sshcopyid_cmd_zip := exec.Command("curl", "-L", sshcopyid_zip_url, "-o", "package.zip")
	err = sshcopyid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sshcopyid_bin_url := "https://cdn.openbsd.org/pub/OpenBSD/OpenSSH/portable/openssh-9.8p1.bin"
	sshcopyid_cmd_bin := exec.Command("curl", "-L", sshcopyid_bin_url, "-o", "binary.bin")
	err = sshcopyid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sshcopyid_src_url := "https://cdn.openbsd.org/pub/OpenBSD/OpenSSH/portable/openssh-9.8p1.src.tar.gz"
	sshcopyid_cmd_src := exec.Command("curl", "-L", sshcopyid_src_url, "-o", "source.tar.gz")
	err = sshcopyid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sshcopyid_cmd_direct := exec.Command("./binary")
	err = sshcopyid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
