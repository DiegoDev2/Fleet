package main

// Rsnapshot - File system snapshot utility (based on rsync)
// Homepage: https://www.rsnapshot.org/

import (
	"fmt"
	
	"os/exec"
)

func installRsnapshot() {
	// Método 1: Descargar y extraer .tar.gz
	rsnapshot_tar_url := "https://github.com/rsnapshot/rsnapshot/releases/download/1.4.5/rsnapshot-1.4.5.tar.gz"
	rsnapshot_cmd_tar := exec.Command("curl", "-L", rsnapshot_tar_url, "-o", "package.tar.gz")
	err := rsnapshot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rsnapshot_zip_url := "https://github.com/rsnapshot/rsnapshot/releases/download/1.4.5/rsnapshot-1.4.5.zip"
	rsnapshot_cmd_zip := exec.Command("curl", "-L", rsnapshot_zip_url, "-o", "package.zip")
	err = rsnapshot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rsnapshot_bin_url := "https://github.com/rsnapshot/rsnapshot/releases/download/1.4.5/rsnapshot-1.4.5.bin"
	rsnapshot_cmd_bin := exec.Command("curl", "-L", rsnapshot_bin_url, "-o", "binary.bin")
	err = rsnapshot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rsnapshot_src_url := "https://github.com/rsnapshot/rsnapshot/releases/download/1.4.5/rsnapshot-1.4.5.src.tar.gz"
	rsnapshot_cmd_src := exec.Command("curl", "-L", rsnapshot_src_url, "-o", "source.tar.gz")
	err = rsnapshot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rsnapshot_cmd_direct := exec.Command("./binary")
	err = rsnapshot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
