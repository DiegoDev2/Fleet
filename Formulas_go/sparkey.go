package main

// Sparkey - Constant key-value store, best for frequent read/infrequent write uses
// Homepage: https://github.com/spotify/sparkey/

import (
	"fmt"
	
	"os/exec"
)

func installSparkey() {
	// Método 1: Descargar y extraer .tar.gz
	sparkey_tar_url := "https://github.com/spotify/sparkey/archive/refs/tags/sparkey-1.0.0.tar.gz"
	sparkey_cmd_tar := exec.Command("curl", "-L", sparkey_tar_url, "-o", "package.tar.gz")
	err := sparkey_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sparkey_zip_url := "https://github.com/spotify/sparkey/archive/refs/tags/sparkey-1.0.0.zip"
	sparkey_cmd_zip := exec.Command("curl", "-L", sparkey_zip_url, "-o", "package.zip")
	err = sparkey_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sparkey_bin_url := "https://github.com/spotify/sparkey/archive/refs/tags/sparkey-1.0.0.bin"
	sparkey_cmd_bin := exec.Command("curl", "-L", sparkey_bin_url, "-o", "binary.bin")
	err = sparkey_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sparkey_src_url := "https://github.com/spotify/sparkey/archive/refs/tags/sparkey-1.0.0.src.tar.gz"
	sparkey_cmd_src := exec.Command("curl", "-L", sparkey_src_url, "-o", "source.tar.gz")
	err = sparkey_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sparkey_cmd_direct := exec.Command("./binary")
	err = sparkey_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: snappy")
exec.Command("latte", "install", "snappy")
}
