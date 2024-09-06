package main

// Tmate - Instant terminal sharing
// Homepage: https://tmate.io/

import (
	"fmt"
	
	"os/exec"
)

func installTmate() {
	// Método 1: Descargar y extraer .tar.gz
	tmate_tar_url := "https://github.com/tmate-io/tmate/archive/refs/tags/2.4.0.tar.gz"
	tmate_cmd_tar := exec.Command("curl", "-L", tmate_tar_url, "-o", "package.tar.gz")
	err := tmate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tmate_zip_url := "https://github.com/tmate-io/tmate/archive/refs/tags/2.4.0.zip"
	tmate_cmd_zip := exec.Command("curl", "-L", tmate_zip_url, "-o", "package.zip")
	err = tmate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tmate_bin_url := "https://github.com/tmate-io/tmate/archive/refs/tags/2.4.0.bin"
	tmate_cmd_bin := exec.Command("curl", "-L", tmate_bin_url, "-o", "binary.bin")
	err = tmate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tmate_src_url := "https://github.com/tmate-io/tmate/archive/refs/tags/2.4.0.src.tar.gz"
	tmate_cmd_src := exec.Command("curl", "-L", tmate_src_url, "-o", "source.tar.gz")
	err = tmate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tmate_cmd_direct := exec.Command("./binary")
	err = tmate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libssh")
	exec.Command("latte", "install", "libssh").Run()
	fmt.Println("Instalando dependencia: msgpack")
	exec.Command("latte", "install", "msgpack").Run()
}
