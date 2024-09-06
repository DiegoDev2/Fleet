package main

// AardvarkShellUtils - Utilities to aid shell scripts or command-line users
// Homepage: http://www.laffeycomputer.com/shellutils.html

import (
	"fmt"
	
	"os/exec"
)

func installAardvarkShellUtils() {
	// Método 1: Descargar y extraer .tar.gz
	aardvarkshellutils_tar_url := "https://web.archive.org/web/20170106105512/downloads.laffeycomputer.com/current_builds/shellutils/aardvark_shell_utils-1.0.tar.gz"
	aardvarkshellutils_cmd_tar := exec.Command("curl", "-L", aardvarkshellutils_tar_url, "-o", "package.tar.gz")
	err := aardvarkshellutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aardvarkshellutils_zip_url := "https://web.archive.org/web/20170106105512/downloads.laffeycomputer.com/current_builds/shellutils/aardvark_shell_utils-1.0.zip"
	aardvarkshellutils_cmd_zip := exec.Command("curl", "-L", aardvarkshellutils_zip_url, "-o", "package.zip")
	err = aardvarkshellutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aardvarkshellutils_bin_url := "https://web.archive.org/web/20170106105512/downloads.laffeycomputer.com/current_builds/shellutils/aardvark_shell_utils-1.0.bin"
	aardvarkshellutils_cmd_bin := exec.Command("curl", "-L", aardvarkshellutils_bin_url, "-o", "binary.bin")
	err = aardvarkshellutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aardvarkshellutils_src_url := "https://web.archive.org/web/20170106105512/downloads.laffeycomputer.com/current_builds/shellutils/aardvark_shell_utils-1.0.src.tar.gz"
	aardvarkshellutils_cmd_src := exec.Command("curl", "-L", aardvarkshellutils_src_url, "-o", "source.tar.gz")
	err = aardvarkshellutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aardvarkshellutils_cmd_direct := exec.Command("./binary")
	err = aardvarkshellutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
