package main

// Md5sha1sum - Hash utilities
// Homepage: http://microbrew.org/tools/md5sha1sum/

import (
	"fmt"
	
	"os/exec"
)

func installMd5sha1sum() {
	// Método 1: Descargar y extraer .tar.gz
	md5sha1sum_tar_url := "http://microbrew.org/tools/md5sha1sum/md5sha1sum-0.9.5.tar.gz"
	md5sha1sum_cmd_tar := exec.Command("curl", "-L", md5sha1sum_tar_url, "-o", "package.tar.gz")
	err := md5sha1sum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	md5sha1sum_zip_url := "http://microbrew.org/tools/md5sha1sum/md5sha1sum-0.9.5.zip"
	md5sha1sum_cmd_zip := exec.Command("curl", "-L", md5sha1sum_zip_url, "-o", "package.zip")
	err = md5sha1sum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	md5sha1sum_bin_url := "http://microbrew.org/tools/md5sha1sum/md5sha1sum-0.9.5.bin"
	md5sha1sum_cmd_bin := exec.Command("curl", "-L", md5sha1sum_bin_url, "-o", "binary.bin")
	err = md5sha1sum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	md5sha1sum_src_url := "http://microbrew.org/tools/md5sha1sum/md5sha1sum-0.9.5.src.tar.gz"
	md5sha1sum_cmd_src := exec.Command("curl", "-L", md5sha1sum_src_url, "-o", "source.tar.gz")
	err = md5sha1sum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	md5sha1sum_cmd_direct := exec.Command("./binary")
	err = md5sha1sum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
