package main

// Mplayershell - Improved visual experience for MPlayer on macOS
// Homepage: https://github.com/lisamelton/MPlayerShell

import (
	"fmt"
	
	"os/exec"
)

func installMplayershell() {
	// Método 1: Descargar y extraer .tar.gz
	mplayershell_tar_url := "https://github.com/lisamelton/MPlayerShell/archive/refs/tags/0.9.3.tar.gz"
	mplayershell_cmd_tar := exec.Command("curl", "-L", mplayershell_tar_url, "-o", "package.tar.gz")
	err := mplayershell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mplayershell_zip_url := "https://github.com/lisamelton/MPlayerShell/archive/refs/tags/0.9.3.zip"
	mplayershell_cmd_zip := exec.Command("curl", "-L", mplayershell_zip_url, "-o", "package.zip")
	err = mplayershell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mplayershell_bin_url := "https://github.com/lisamelton/MPlayerShell/archive/refs/tags/0.9.3.bin"
	mplayershell_cmd_bin := exec.Command("curl", "-L", mplayershell_bin_url, "-o", "binary.bin")
	err = mplayershell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mplayershell_src_url := "https://github.com/lisamelton/MPlayerShell/archive/refs/tags/0.9.3.src.tar.gz"
	mplayershell_cmd_src := exec.Command("curl", "-L", mplayershell_src_url, "-o", "source.tar.gz")
	err = mplayershell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mplayershell_cmd_direct := exec.Command("./binary")
	err = mplayershell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mplayer")
exec.Command("latte", "install", "mplayer")
}
