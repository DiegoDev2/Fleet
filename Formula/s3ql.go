package main

// S3ql - POSIX-compliant FUSE filesystem using object store as block storage
// Homepage: https://github.com/s3ql/s3ql

import (
	"fmt"
	
	"os/exec"
)

func installS3ql() {
	// Método 1: Descargar y extraer .tar.gz
	s3ql_tar_url := "https://github.com/s3ql/s3ql/releases/download/s3ql-5.2.2/s3ql-5.2.2.tar.gz"
	s3ql_cmd_tar := exec.Command("curl", "-L", s3ql_tar_url, "-o", "package.tar.gz")
	err := s3ql_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	s3ql_zip_url := "https://github.com/s3ql/s3ql/releases/download/s3ql-5.2.2/s3ql-5.2.2.zip"
	s3ql_cmd_zip := exec.Command("curl", "-L", s3ql_zip_url, "-o", "package.zip")
	err = s3ql_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	s3ql_bin_url := "https://github.com/s3ql/s3ql/releases/download/s3ql-5.2.2/s3ql-5.2.2.bin"
	s3ql_cmd_bin := exec.Command("curl", "-L", s3ql_bin_url, "-o", "binary.bin")
	err = s3ql_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	s3ql_src_url := "https://github.com/s3ql/s3ql/releases/download/s3ql-5.2.2/s3ql-5.2.2.src.tar.gz"
	s3ql_cmd_src := exec.Command("curl", "-L", s3ql_src_url, "-o", "source.tar.gz")
	err = s3ql_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	s3ql_cmd_direct := exec.Command("./binary")
	err = s3ql_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cython")
	exec.Command("latte", "install", "cython").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libffi")
	exec.Command("latte", "install", "libffi").Run()
	fmt.Println("Instalando dependencia: libfuse")
	exec.Command("latte", "install", "libfuse").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
