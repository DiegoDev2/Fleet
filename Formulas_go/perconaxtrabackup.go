package main

// PerconaXtrabackup - Open source hot backup tool for InnoDB and XtraDB databases
// Homepage: https://www.percona.com/software/mysql-database/percona-xtrabackup

import (
	"fmt"
	
	"os/exec"
)

func installPerconaXtrabackup() {
	// Método 1: Descargar y extraer .tar.gz
	perconaxtrabackup_tar_url := "https://downloads.percona.com/downloads/Percona-XtraBackup-LATEST/Percona-XtraBackup-8.0.35-31/source/tarball/percona-xtrabackup-8.0.35-31.tar.gz"
	perconaxtrabackup_cmd_tar := exec.Command("curl", "-L", perconaxtrabackup_tar_url, "-o", "package.tar.gz")
	err := perconaxtrabackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	perconaxtrabackup_zip_url := "https://downloads.percona.com/downloads/Percona-XtraBackup-LATEST/Percona-XtraBackup-8.0.35-31/source/tarball/percona-xtrabackup-8.0.35-31.zip"
	perconaxtrabackup_cmd_zip := exec.Command("curl", "-L", perconaxtrabackup_zip_url, "-o", "package.zip")
	err = perconaxtrabackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	perconaxtrabackup_bin_url := "https://downloads.percona.com/downloads/Percona-XtraBackup-LATEST/Percona-XtraBackup-8.0.35-31/source/tarball/percona-xtrabackup-8.0.35-31.bin"
	perconaxtrabackup_cmd_bin := exec.Command("curl", "-L", perconaxtrabackup_bin_url, "-o", "binary.bin")
	err = perconaxtrabackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	perconaxtrabackup_src_url := "https://downloads.percona.com/downloads/Percona-XtraBackup-LATEST/Percona-XtraBackup-8.0.35-31/source/tarball/percona-xtrabackup-8.0.35-31.src.tar.gz"
	perconaxtrabackup_cmd_src := exec.Command("curl", "-L", perconaxtrabackup_src_url, "-o", "source.tar.gz")
	err = perconaxtrabackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	perconaxtrabackup_cmd_direct := exec.Command("./binary")
	err = perconaxtrabackup_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: sphinx-doc")
exec.Command("latte", "install", "sphinx-doc")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: libev")
exec.Command("latte", "install", "libev")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: libfido2")
exec.Command("latte", "install", "libfido2")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: mysql-client")
exec.Command("latte", "install", "mysql-client")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: protobuf@21")
exec.Command("latte", "install", "protobuf@21")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: patchelf")
exec.Command("latte", "install", "patchelf")
	fmt.Println("Instalando dependencia: libaio")
exec.Command("latte", "install", "libaio")
	fmt.Println("Instalando dependencia: procps")
exec.Command("latte", "install", "procps")
}
