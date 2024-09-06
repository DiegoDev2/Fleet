package main

// FcitxRemoteForOsx - Handle input method in command-line
// Homepage: https://github.com/xcodebuild/fcitx-remote-for-osx

import (
	"fmt"
	
	"os/exec"
)

func installFcitxRemoteForOsx() {
	// Método 1: Descargar y extraer .tar.gz
	fcitxremoteforosx_tar_url := "https://github.com/xcodebuild/fcitx-remote-for-osx/archive/refs/tags/0.4.0.tar.gz"
	fcitxremoteforosx_cmd_tar := exec.Command("curl", "-L", fcitxremoteforosx_tar_url, "-o", "package.tar.gz")
	err := fcitxremoteforosx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fcitxremoteforosx_zip_url := "https://github.com/xcodebuild/fcitx-remote-for-osx/archive/refs/tags/0.4.0.zip"
	fcitxremoteforosx_cmd_zip := exec.Command("curl", "-L", fcitxremoteforosx_zip_url, "-o", "package.zip")
	err = fcitxremoteforosx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fcitxremoteforosx_bin_url := "https://github.com/xcodebuild/fcitx-remote-for-osx/archive/refs/tags/0.4.0.bin"
	fcitxremoteforosx_cmd_bin := exec.Command("curl", "-L", fcitxremoteforosx_bin_url, "-o", "binary.bin")
	err = fcitxremoteforosx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fcitxremoteforosx_src_url := "https://github.com/xcodebuild/fcitx-remote-for-osx/archive/refs/tags/0.4.0.src.tar.gz"
	fcitxremoteforosx_cmd_src := exec.Command("curl", "-L", fcitxremoteforosx_src_url, "-o", "source.tar.gz")
	err = fcitxremoteforosx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fcitxremoteforosx_cmd_direct := exec.Command("./binary")
	err = fcitxremoteforosx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
