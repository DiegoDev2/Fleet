package main

// Proto - Pluggable multi-language version manager
// Homepage: https://moonrepo.dev/proto

import (
	"fmt"
	
	"os/exec"
)

func installProto() {
	// Método 1: Descargar y extraer .tar.gz
	proto_tar_url := "https://github.com/moonrepo/proto/archive/refs/tags/v0.40.4.tar.gz"
	proto_cmd_tar := exec.Command("curl", "-L", proto_tar_url, "-o", "package.tar.gz")
	err := proto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proto_zip_url := "https://github.com/moonrepo/proto/archive/refs/tags/v0.40.4.zip"
	proto_cmd_zip := exec.Command("curl", "-L", proto_zip_url, "-o", "package.zip")
	err = proto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proto_bin_url := "https://github.com/moonrepo/proto/archive/refs/tags/v0.40.4.bin"
	proto_cmd_bin := exec.Command("curl", "-L", proto_bin_url, "-o", "binary.bin")
	err = proto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proto_src_url := "https://github.com/moonrepo/proto/archive/refs/tags/v0.40.4.src.tar.gz"
	proto_cmd_src := exec.Command("curl", "-L", proto_src_url, "-o", "source.tar.gz")
	err = proto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proto_cmd_direct := exec.Command("./binary")
	err = proto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
