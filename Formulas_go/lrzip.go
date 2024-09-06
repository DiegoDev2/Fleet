package main

// Lrzip - Compression program with a very high compression ratio
// Homepage: http://lrzip.kolivas.org

import (
	"fmt"
	
	"os/exec"
)

func installLrzip() {
	// Método 1: Descargar y extraer .tar.gz
	lrzip_tar_url := "http://ck.kolivas.org/apps/lrzip/lrzip-0.641.tar.xz"
	lrzip_cmd_tar := exec.Command("curl", "-L", lrzip_tar_url, "-o", "package.tar.gz")
	err := lrzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lrzip_zip_url := "http://ck.kolivas.org/apps/lrzip/lrzip-0.641.tar.xz"
	lrzip_cmd_zip := exec.Command("curl", "-L", lrzip_zip_url, "-o", "package.zip")
	err = lrzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lrzip_bin_url := "http://ck.kolivas.org/apps/lrzip/lrzip-0.641.tar.xz"
	lrzip_cmd_bin := exec.Command("curl", "-L", lrzip_bin_url, "-o", "binary.bin")
	err = lrzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lrzip_src_url := "http://ck.kolivas.org/apps/lrzip/lrzip-0.641.tar.xz"
	lrzip_cmd_src := exec.Command("curl", "-L", lrzip_src_url, "-o", "source.tar.gz")
	err = lrzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lrzip_cmd_direct := exec.Command("./binary")
	err = lrzip_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: lzo")
exec.Command("latte", "install", "lzo")
	fmt.Println("Instalando dependencia: nasm")
exec.Command("latte", "install", "nasm")
}
