package main

// Argon2 - Password hashing library and CLI utility
// Homepage: https://github.com/P-H-C/phc-winner-argon2

import (
	"fmt"
	
	"os/exec"
)

func installArgon2() {
	// Método 1: Descargar y extraer .tar.gz
	argon2_tar_url := "https://github.com/P-H-C/phc-winner-argon2/archive/refs/tags/20190702.tar.gz"
	argon2_cmd_tar := exec.Command("curl", "-L", argon2_tar_url, "-o", "package.tar.gz")
	err := argon2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	argon2_zip_url := "https://github.com/P-H-C/phc-winner-argon2/archive/refs/tags/20190702.zip"
	argon2_cmd_zip := exec.Command("curl", "-L", argon2_zip_url, "-o", "package.zip")
	err = argon2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	argon2_bin_url := "https://github.com/P-H-C/phc-winner-argon2/archive/refs/tags/20190702.bin"
	argon2_cmd_bin := exec.Command("curl", "-L", argon2_bin_url, "-o", "binary.bin")
	err = argon2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	argon2_src_url := "https://github.com/P-H-C/phc-winner-argon2/archive/refs/tags/20190702.src.tar.gz"
	argon2_cmd_src := exec.Command("curl", "-L", argon2_src_url, "-o", "source.tar.gz")
	err = argon2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	argon2_cmd_direct := exec.Command("./binary")
	err = argon2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
