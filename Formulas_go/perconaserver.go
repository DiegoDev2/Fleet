package main

// PerconaServer - Drop-in MySQL replacement
// Homepage: https://www.percona.com

import (
	"fmt"
	
	"os/exec"
)

func installPerconaServer() {
	// Método 1: Descargar y extraer .tar.gz
	perconaserver_tar_url := "https://downloads.percona.com/downloads/Percona-Server-8.0/Percona-Server-8.0.36-28/source/tarball/percona-server-8.0.36-28.tar.gz"
	perconaserver_cmd_tar := exec.Command("curl", "-L", perconaserver_tar_url, "-o", "package.tar.gz")
	err := perconaserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	perconaserver_zip_url := "https://downloads.percona.com/downloads/Percona-Server-8.0/Percona-Server-8.0.36-28/source/tarball/percona-server-8.0.36-28.zip"
	perconaserver_cmd_zip := exec.Command("curl", "-L", perconaserver_zip_url, "-o", "package.zip")
	err = perconaserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	perconaserver_bin_url := "https://downloads.percona.com/downloads/Percona-Server-8.0/Percona-Server-8.0.36-28/source/tarball/percona-server-8.0.36-28.bin"
	perconaserver_cmd_bin := exec.Command("curl", "-L", perconaserver_bin_url, "-o", "binary.bin")
	err = perconaserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	perconaserver_src_url := "https://downloads.percona.com/downloads/Percona-Server-8.0/Percona-Server-8.0.36-28/source/tarball/percona-server-8.0.36-28.src.tar.gz"
	perconaserver_cmd_src := exec.Command("curl", "-L", perconaserver_src_url, "-o", "source.tar.gz")
	err = perconaserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	perconaserver_cmd_direct := exec.Command("./binary")
	err = perconaserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: libfido2")
exec.Command("latte", "install", "libfido2")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: openldap")
exec.Command("latte", "install", "openldap")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: protobuf@21")
exec.Command("latte", "install", "protobuf@21")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: patchelf")
exec.Command("latte", "install", "patchelf")
	fmt.Println("Instalando dependencia: libtirpc")
exec.Command("latte", "install", "libtirpc")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
