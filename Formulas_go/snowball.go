package main

// Snowball - Stemming algorithms
// Homepage: https://snowballstem.org

import (
	"fmt"
	
	"os/exec"
)

func installSnowball() {
	// Método 1: Descargar y extraer .tar.gz
	snowball_tar_url := "https://github.com/snowballstem/snowball/archive/refs/tags/v2.2.0.tar.gz"
	snowball_cmd_tar := exec.Command("curl", "-L", snowball_tar_url, "-o", "package.tar.gz")
	err := snowball_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snowball_zip_url := "https://github.com/snowballstem/snowball/archive/refs/tags/v2.2.0.zip"
	snowball_cmd_zip := exec.Command("curl", "-L", snowball_zip_url, "-o", "package.zip")
	err = snowball_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snowball_bin_url := "https://github.com/snowballstem/snowball/archive/refs/tags/v2.2.0.bin"
	snowball_cmd_bin := exec.Command("curl", "-L", snowball_bin_url, "-o", "binary.bin")
	err = snowball_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snowball_src_url := "https://github.com/snowballstem/snowball/archive/refs/tags/v2.2.0.src.tar.gz"
	snowball_cmd_src := exec.Command("curl", "-L", snowball_src_url, "-o", "source.tar.gz")
	err = snowball_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snowball_cmd_direct := exec.Command("./binary")
	err = snowball_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
