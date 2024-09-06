package main

// S3fs - FUSE-based file system backed by Amazon S3
// Homepage: https://github.com/s3fs-fuse/s3fs-fuse/wiki

import (
	"fmt"
	
	"os/exec"
)

func installS3fs() {
	// Método 1: Descargar y extraer .tar.gz
	s3fs_tar_url := "https://github.com/s3fs-fuse/s3fs-fuse/archive/refs/tags/v1.94.tar.gz"
	s3fs_cmd_tar := exec.Command("curl", "-L", s3fs_tar_url, "-o", "package.tar.gz")
	err := s3fs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	s3fs_zip_url := "https://github.com/s3fs-fuse/s3fs-fuse/archive/refs/tags/v1.94.zip"
	s3fs_cmd_zip := exec.Command("curl", "-L", s3fs_zip_url, "-o", "package.zip")
	err = s3fs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	s3fs_bin_url := "https://github.com/s3fs-fuse/s3fs-fuse/archive/refs/tags/v1.94.bin"
	s3fs_cmd_bin := exec.Command("curl", "-L", s3fs_bin_url, "-o", "binary.bin")
	err = s3fs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	s3fs_src_url := "https://github.com/s3fs-fuse/s3fs-fuse/archive/refs/tags/v1.94.src.tar.gz"
	s3fs_cmd_src := exec.Command("curl", "-L", s3fs_src_url, "-o", "source.tar.gz")
	err = s3fs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	s3fs_cmd_direct := exec.Command("./binary")
	err = s3fs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: libfuse@2")
	exec.Command("latte", "install", "libfuse@2").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: nettle")
	exec.Command("latte", "install", "nettle").Run()
}
