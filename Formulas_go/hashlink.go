package main

// Hashlink - Virtual machine for Haxe
// Homepage: https://hashlink.haxe.org/

import (
	"fmt"
	
	"os/exec"
)

func installHashlink() {
	// Método 1: Descargar y extraer .tar.gz
	hashlink_tar_url := "https://github.com/HaxeFoundation/hashlink/archive/refs/tags/1.14.tar.gz"
	hashlink_cmd_tar := exec.Command("curl", "-L", hashlink_tar_url, "-o", "package.tar.gz")
	err := hashlink_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hashlink_zip_url := "https://github.com/HaxeFoundation/hashlink/archive/refs/tags/1.14.zip"
	hashlink_cmd_zip := exec.Command("curl", "-L", hashlink_zip_url, "-o", "package.zip")
	err = hashlink_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hashlink_bin_url := "https://github.com/HaxeFoundation/hashlink/archive/refs/tags/1.14.bin"
	hashlink_cmd_bin := exec.Command("curl", "-L", hashlink_bin_url, "-o", "binary.bin")
	err = hashlink_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hashlink_src_url := "https://github.com/HaxeFoundation/hashlink/archive/refs/tags/1.14.src.tar.gz"
	hashlink_cmd_src := exec.Command("curl", "-L", hashlink_src_url, "-o", "source.tar.gz")
	err = hashlink_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hashlink_cmd_direct := exec.Command("./binary")
	err = hashlink_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: haxe")
exec.Command("latte", "install", "haxe")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libuv")
exec.Command("latte", "install", "libuv")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: mbedtls")
exec.Command("latte", "install", "mbedtls")
	fmt.Println("Instalando dependencia: openal-soft")
exec.Command("latte", "install", "openal-soft")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
}
