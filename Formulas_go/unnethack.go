package main

// Unnethack - Fork of Nethack
// Homepage: https://unnethack.wordpress.com/

import (
	"fmt"
	
	"os/exec"
)

func installUnnethack() {
	// Método 1: Descargar y extraer .tar.gz
	unnethack_tar_url := "https://github.com/UnNetHack/UnNetHack/archive/refs/tags/5.3.2.tar.gz"
	unnethack_cmd_tar := exec.Command("curl", "-L", unnethack_tar_url, "-o", "package.tar.gz")
	err := unnethack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unnethack_zip_url := "https://github.com/UnNetHack/UnNetHack/archive/refs/tags/5.3.2.zip"
	unnethack_cmd_zip := exec.Command("curl", "-L", unnethack_zip_url, "-o", "package.zip")
	err = unnethack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unnethack_bin_url := "https://github.com/UnNetHack/UnNetHack/archive/refs/tags/5.3.2.bin"
	unnethack_cmd_bin := exec.Command("curl", "-L", unnethack_bin_url, "-o", "binary.bin")
	err = unnethack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unnethack_src_url := "https://github.com/UnNetHack/UnNetHack/archive/refs/tags/5.3.2.src.tar.gz"
	unnethack_cmd_src := exec.Command("curl", "-L", unnethack_src_url, "-o", "source.tar.gz")
	err = unnethack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unnethack_cmd_direct := exec.Command("./binary")
	err = unnethack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
