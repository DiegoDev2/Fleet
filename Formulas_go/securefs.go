package main

// Securefs - Filesystem with transparent authenticated encryption
// Homepage: https://github.com/netheril96/securefs

import (
	"fmt"
	
	"os/exec"
)

func installSecurefs() {
	// Método 1: Descargar y extraer .tar.gz
	securefs_tar_url := "https://github.com/netheril96/securefs/archive/refs/tags/v1.0.0.tar.gz"
	securefs_cmd_tar := exec.Command("curl", "-L", securefs_tar_url, "-o", "package.tar.gz")
	err := securefs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	securefs_zip_url := "https://github.com/netheril96/securefs/archive/refs/tags/v1.0.0.zip"
	securefs_cmd_zip := exec.Command("curl", "-L", securefs_zip_url, "-o", "package.zip")
	err = securefs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	securefs_bin_url := "https://github.com/netheril96/securefs/archive/refs/tags/v1.0.0.bin"
	securefs_cmd_bin := exec.Command("curl", "-L", securefs_bin_url, "-o", "binary.bin")
	err = securefs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	securefs_src_url := "https://github.com/netheril96/securefs/archive/refs/tags/v1.0.0.src.tar.gz"
	securefs_cmd_src := exec.Command("curl", "-L", securefs_src_url, "-o", "source.tar.gz")
	err = securefs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	securefs_cmd_direct := exec.Command("./binary")
	err = securefs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: tclap")
exec.Command("latte", "install", "tclap")
	fmt.Println("Instalando dependencia: abseil")
exec.Command("latte", "install", "abseil")
	fmt.Println("Instalando dependencia: argon2")
exec.Command("latte", "install", "argon2")
	fmt.Println("Instalando dependencia: cryptopp")
exec.Command("latte", "install", "cryptopp")
	fmt.Println("Instalando dependencia: fruit")
exec.Command("latte", "install", "fruit")
	fmt.Println("Instalando dependencia: jsoncpp")
exec.Command("latte", "install", "jsoncpp")
	fmt.Println("Instalando dependencia: libfuse@2")
exec.Command("latte", "install", "libfuse@2")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: uni-algo")
exec.Command("latte", "install", "uni-algo")
	fmt.Println("Instalando dependencia: utf8proc")
exec.Command("latte", "install", "utf8proc")
}
