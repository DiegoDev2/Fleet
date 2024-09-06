package main

// Zzuf - Transparent application input fuzzer
// Homepage: http://caca.zoy.org/wiki/zzuf

import (
	"fmt"
	
	"os/exec"
)

func installZzuf() {
	// Método 1: Descargar y extraer .tar.gz
	zzuf_tar_url := "https://github.com/samhocevar/zzuf/releases/download/v0.15/zzuf-0.15.tar.bz2"
	zzuf_cmd_tar := exec.Command("curl", "-L", zzuf_tar_url, "-o", "package.tar.gz")
	err := zzuf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zzuf_zip_url := "https://github.com/samhocevar/zzuf/releases/download/v0.15/zzuf-0.15.tar.bz2"
	zzuf_cmd_zip := exec.Command("curl", "-L", zzuf_zip_url, "-o", "package.zip")
	err = zzuf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zzuf_bin_url := "https://github.com/samhocevar/zzuf/releases/download/v0.15/zzuf-0.15.tar.bz2"
	zzuf_cmd_bin := exec.Command("curl", "-L", zzuf_bin_url, "-o", "binary.bin")
	err = zzuf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zzuf_src_url := "https://github.com/samhocevar/zzuf/releases/download/v0.15/zzuf-0.15.tar.bz2"
	zzuf_cmd_src := exec.Command("curl", "-L", zzuf_src_url, "-o", "source.tar.gz")
	err = zzuf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zzuf_cmd_direct := exec.Command("./binary")
	err = zzuf_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
