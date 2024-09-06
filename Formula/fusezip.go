package main

// FuseZip - FUSE file system to create & manipulate ZIP archives
// Homepage: https://bitbucket.org/agalanin/fuse-zip

import (
	"fmt"
	
	"os/exec"
)

func installFuseZip() {
	// Método 1: Descargar y extraer .tar.gz
	fusezip_tar_url := "https://bitbucket.org/agalanin/fuse-zip/downloads/fuse-zip-0.7.2.tar.gz"
	fusezip_cmd_tar := exec.Command("curl", "-L", fusezip_tar_url, "-o", "package.tar.gz")
	err := fusezip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fusezip_zip_url := "https://bitbucket.org/agalanin/fuse-zip/downloads/fuse-zip-0.7.2.zip"
	fusezip_cmd_zip := exec.Command("curl", "-L", fusezip_zip_url, "-o", "package.zip")
	err = fusezip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fusezip_bin_url := "https://bitbucket.org/agalanin/fuse-zip/downloads/fuse-zip-0.7.2.bin"
	fusezip_cmd_bin := exec.Command("curl", "-L", fusezip_bin_url, "-o", "binary.bin")
	err = fusezip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fusezip_src_url := "https://bitbucket.org/agalanin/fuse-zip/downloads/fuse-zip-0.7.2.src.tar.gz"
	fusezip_cmd_src := exec.Command("curl", "-L", fusezip_src_url, "-o", "source.tar.gz")
	err = fusezip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fusezip_cmd_direct := exec.Command("./binary")
	err = fusezip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libfuse@2")
	exec.Command("latte", "install", "libfuse@2").Run()
	fmt.Println("Instalando dependencia: libzip")
	exec.Command("latte", "install", "libzip").Run()
}
