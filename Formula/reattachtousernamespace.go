package main

// ReattachToUserNamespace - Reattach process (e.g., tmux) to background
// Homepage: https://github.com/ChrisJohnsen/tmux-MacOSX-pasteboard

import (
	"fmt"
	
	"os/exec"
)

func installReattachToUserNamespace() {
	// Método 1: Descargar y extraer .tar.gz
	reattachtousernamespace_tar_url := "https://github.com/ChrisJohnsen/tmux-MacOSX-pasteboard/archive/refs/tags/v2.9.tar.gz"
	reattachtousernamespace_cmd_tar := exec.Command("curl", "-L", reattachtousernamespace_tar_url, "-o", "package.tar.gz")
	err := reattachtousernamespace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reattachtousernamespace_zip_url := "https://github.com/ChrisJohnsen/tmux-MacOSX-pasteboard/archive/refs/tags/v2.9.zip"
	reattachtousernamespace_cmd_zip := exec.Command("curl", "-L", reattachtousernamespace_zip_url, "-o", "package.zip")
	err = reattachtousernamespace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reattachtousernamespace_bin_url := "https://github.com/ChrisJohnsen/tmux-MacOSX-pasteboard/archive/refs/tags/v2.9.bin"
	reattachtousernamespace_cmd_bin := exec.Command("curl", "-L", reattachtousernamespace_bin_url, "-o", "binary.bin")
	err = reattachtousernamespace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reattachtousernamespace_src_url := "https://github.com/ChrisJohnsen/tmux-MacOSX-pasteboard/archive/refs/tags/v2.9.src.tar.gz"
	reattachtousernamespace_cmd_src := exec.Command("curl", "-L", reattachtousernamespace_src_url, "-o", "source.tar.gz")
	err = reattachtousernamespace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reattachtousernamespace_cmd_direct := exec.Command("./binary")
	err = reattachtousernamespace_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
