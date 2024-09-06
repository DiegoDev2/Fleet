package main

// Kstart - Modified version of kinit that can use keytabs to authenticate
// Homepage: https://www.eyrie.org/~eagle/software/kstart/

import (
	"fmt"
	
	"os/exec"
)

func installKstart() {
	// Método 1: Descargar y extraer .tar.gz
	kstart_tar_url := "https://archives.eyrie.org/software/kerberos/kstart-4.3.tar.xz"
	kstart_cmd_tar := exec.Command("curl", "-L", kstart_tar_url, "-o", "package.tar.gz")
	err := kstart_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kstart_zip_url := "https://archives.eyrie.org/software/kerberos/kstart-4.3.tar.xz"
	kstart_cmd_zip := exec.Command("curl", "-L", kstart_zip_url, "-o", "package.zip")
	err = kstart_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kstart_bin_url := "https://archives.eyrie.org/software/kerberos/kstart-4.3.tar.xz"
	kstart_cmd_bin := exec.Command("curl", "-L", kstart_bin_url, "-o", "binary.bin")
	err = kstart_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kstart_src_url := "https://archives.eyrie.org/software/kerberos/kstart-4.3.tar.xz"
	kstart_cmd_src := exec.Command("curl", "-L", kstart_src_url, "-o", "source.tar.gz")
	err = kstart_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kstart_cmd_direct := exec.Command("./binary")
	err = kstart_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
