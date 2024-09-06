package main

// Sshpass - Non-interactive SSH password auth
// Homepage: https://sourceforge.net/projects/sshpass/

import (
	"fmt"
	
	"os/exec"
)

func installSshpass() {
	// Método 1: Descargar y extraer .tar.gz
	sshpass_tar_url := "https://master.dl.sourceforge.net/project/sshpass/sshpass/1.10/sshpass-1.10.tar.gz"
	sshpass_cmd_tar := exec.Command("curl", "-L", sshpass_tar_url, "-o", "package.tar.gz")
	err := sshpass_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sshpass_zip_url := "https://master.dl.sourceforge.net/project/sshpass/sshpass/1.10/sshpass-1.10.zip"
	sshpass_cmd_zip := exec.Command("curl", "-L", sshpass_zip_url, "-o", "package.zip")
	err = sshpass_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sshpass_bin_url := "https://master.dl.sourceforge.net/project/sshpass/sshpass/1.10/sshpass-1.10.bin"
	sshpass_cmd_bin := exec.Command("curl", "-L", sshpass_bin_url, "-o", "binary.bin")
	err = sshpass_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sshpass_src_url := "https://master.dl.sourceforge.net/project/sshpass/sshpass/1.10/sshpass-1.10.src.tar.gz"
	sshpass_cmd_src := exec.Command("curl", "-L", sshpass_src_url, "-o", "source.tar.gz")
	err = sshpass_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sshpass_cmd_direct := exec.Command("./binary")
	err = sshpass_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
