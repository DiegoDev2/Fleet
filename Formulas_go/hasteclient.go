package main

// HasteClient - CLI client for haste-server
// Homepage: https://hastebin.com/

import (
	"fmt"
	
	"os/exec"
)

func installHasteClient() {
	// Método 1: Descargar y extraer .tar.gz
	hasteclient_tar_url := "https://github.com/toptal/haste-client/archive/refs/tags/v0.3.0.tar.gz"
	hasteclient_cmd_tar := exec.Command("curl", "-L", hasteclient_tar_url, "-o", "package.tar.gz")
	err := hasteclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hasteclient_zip_url := "https://github.com/toptal/haste-client/archive/refs/tags/v0.3.0.zip"
	hasteclient_cmd_zip := exec.Command("curl", "-L", hasteclient_zip_url, "-o", "package.zip")
	err = hasteclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hasteclient_bin_url := "https://github.com/toptal/haste-client/archive/refs/tags/v0.3.0.bin"
	hasteclient_cmd_bin := exec.Command("curl", "-L", hasteclient_bin_url, "-o", "binary.bin")
	err = hasteclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hasteclient_src_url := "https://github.com/toptal/haste-client/archive/refs/tags/v0.3.0.src.tar.gz"
	hasteclient_cmd_src := exec.Command("curl", "-L", hasteclient_src_url, "-o", "source.tar.gz")
	err = hasteclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hasteclient_cmd_direct := exec.Command("./binary")
	err = hasteclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
