import 'reflect-metadata'
import express from "express";
import { Application } from 'express';
import { Server } from 'http';
import { Container } from 'inversify';
import { IInstaller } from "./contracts/IInstaller";
import { MiddlewareInstaller } from './installers/MiddlewareInstaller';


const AppContainer = new Container();

// AppContainer
AppContainer.bind<Container>("AppContainer").toConstantValue(AppContainer);

// Application and Server
AppContainer.bind<Application>("Application").toConstantValue(express());
AppContainer.bind<Server>(Server).toConstantValue(new Server(AppContainer.get<Application>("Application") as any));

// Installers
AppContainer.bind<IInstaller>("IInstaller").to(MiddlewareInstaller);

export { AppContainer };