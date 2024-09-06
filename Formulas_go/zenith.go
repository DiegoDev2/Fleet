package main

// Zenith - In terminal graphical metrics for your *nix system
// Homepage: https://github.com/bvaisvil/zenith/

import (
	"fmt"
	
	"os/exec"
)

func installZenith() {
	// Método 1: Descargar y extraer .tar.gz
	zenith_tar_url := "https://github.com/bvaisvil/zenith/archive/refs/tags/0.14.1.tar.gz"
	zenith_cmd_tar := exec.Command("curl", "-L", zenith_tar_url, "-o", "package.tar.gz")
	err := zenith_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zenith_zip_url := "https://github.com/bvaisvil/zenith/archive/refs/tags/0.14.1.zip"
	zenith_cmd_zip := exec.Command("curl", "-L", zenith_zip_url, "-o", "package.zip")
	err = zenith_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zenith_bin_url := "https://github.com/bvaisvil/zenith/archive/refs/tags/0.14.1.bin"
	zenith_cmd_bin := exec.Command("curl", "-L", zenith_bin_url, "-o", "binary.bin")
	err = zenith_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zenith_src_url := "https://github.com/bvaisvil/zenith/archive/refs/tags/0.14.1.src.tar.gz"
	zenith_cmd_src := exec.Command("curl", "-L", zenith_src_url, "-o", "source.tar.gz")
	err = zenith_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zenith_cmd_direct := exec.Command("./binary")
	err = zenith_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
