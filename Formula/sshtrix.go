package main

// Sshtrix - SSH login cracker
// Homepage: https://nullsecurity.net/tools/cracker.html

import (
	"fmt"
	
	"os/exec"
)

func installSshtrix() {
	// Método 1: Descargar y extraer .tar.gz
	sshtrix_tar_url := "https://github.com/nullsecuritynet/tools/raw/main/cracker/sshtrix/release/sshtrix-0.0.3.tar.gz"
	sshtrix_cmd_tar := exec.Command("curl", "-L", sshtrix_tar_url, "-o", "package.tar.gz")
	err := sshtrix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sshtrix_zip_url := "https://github.com/nullsecuritynet/tools/raw/main/cracker/sshtrix/release/sshtrix-0.0.3.zip"
	sshtrix_cmd_zip := exec.Command("curl", "-L", sshtrix_zip_url, "-o", "package.zip")
	err = sshtrix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sshtrix_bin_url := "https://github.com/nullsecuritynet/tools/raw/main/cracker/sshtrix/release/sshtrix-0.0.3.bin"
	sshtrix_cmd_bin := exec.Command("curl", "-L", sshtrix_bin_url, "-o", "binary.bin")
	err = sshtrix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sshtrix_src_url := "https://github.com/nullsecuritynet/tools/raw/main/cracker/sshtrix/release/sshtrix-0.0.3.src.tar.gz"
	sshtrix_cmd_src := exec.Command("curl", "-L", sshtrix_src_url, "-o", "source.tar.gz")
	err = sshtrix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sshtrix_cmd_direct := exec.Command("./binary")
	err = sshtrix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libssh")
	exec.Command("latte", "install", "libssh").Run()
}
