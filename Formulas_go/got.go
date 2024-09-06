package main

// Got - Version control system
// Homepage: https://gameoftrees.org/

import (
	"fmt"
	
	"os/exec"
)

func installGot() {
	// Método 1: Descargar y extraer .tar.gz
	got_tar_url := "https://gameoftrees.org/releases/portable/got-portable-0.102.tar.gz"
	got_cmd_tar := exec.Command("curl", "-L", got_tar_url, "-o", "package.tar.gz")
	err := got_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	got_zip_url := "https://gameoftrees.org/releases/portable/got-portable-0.102.zip"
	got_cmd_zip := exec.Command("curl", "-L", got_zip_url, "-o", "package.zip")
	err = got_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	got_bin_url := "https://gameoftrees.org/releases/portable/got-portable-0.102.bin"
	got_cmd_bin := exec.Command("curl", "-L", got_bin_url, "-o", "binary.bin")
	err = got_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	got_src_url := "https://gameoftrees.org/releases/portable/got-portable-0.102.src.tar.gz"
	got_cmd_src := exec.Command("curl", "-L", got_src_url, "-o", "source.tar.gz")
	err = got_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	got_cmd_direct := exec.Command("./binary")
	err = got_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: libretls")
exec.Command("latte", "install", "libretls")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: libbsd")
exec.Command("latte", "install", "libbsd")
	fmt.Println("Instalando dependencia: libmd")
exec.Command("latte", "install", "libmd")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
